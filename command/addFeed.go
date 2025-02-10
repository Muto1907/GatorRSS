package command

import (
	"context"
	"fmt"
	"time"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *config.State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addFeed <name> <url>")
	}
	currUser, err := s.Db.GetUser(context.Background(), s.Config.CurrentUsername)
	if err != nil {
		return fmt.Errorf("couldn't fetch User to add Feed to: %w", err)
	}
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currUser.ID,
	}
	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't add feed: %w", err)
	}
	fmt.Printf("Feed added successfuly:\n%s", feedToString(feed))
	return nil
}

func feedToString(feed database.Feed) string {
	return fmt.Sprintf("Id: %v,\nName: %s,\nURL: %s\nAssigned to: %s\ncreatedAt: %v,\nupdatedAt: %v\n", feed.ID, feed.Name, feed.Url, feed.UserID, feed.CreatedAt, feed.UpdatedAt)
}
