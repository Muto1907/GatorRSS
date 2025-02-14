package command

import (
	"context"
	"fmt"
	"time"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <feed_url>")
	}
	feed, err := s.Db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	follow, err := s.Db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating follow entry: %w", err)
	}
	fmt.Printf("User %s is now following %s!\n", follow.UsersName, follow.FeedName)
	return nil
}
