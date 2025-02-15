package rss

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
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
	fmt.Printf("Articles of %s\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Printf("%s\n", item.Title)
	}
	fmt.Println("----------------------------------------")
	return nil
}
