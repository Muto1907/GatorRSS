package rss

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
	"github.com/google/uuid"
)

func ScrapeFeeds(s *config.State) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch next feed: %w", err)
	}
	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC()},
		ID:            nextFeed.ID,
		UpdatedAt:     time.Now().UTC(),
	}
	err = s.Db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}
	feed, err := FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	for _, item := range feed.Channel.Item {
		publishedAt := sql.NullTime{}
		pTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return fmt.Errorf("couldn't parse time of the post: %w", err)
		}
		publishedAt.Time = pTime
		publishedAt.Valid = true
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		}
		_, err = s.Db.CreatePost(context.Background(), params)
		if err != nil && err.Error() != `pq: duplicate key value violates unique constraint "posts_url_key"` {
			return fmt.Errorf("couldn't Create Post: %w", err)
		}
	}
	return nil
}
