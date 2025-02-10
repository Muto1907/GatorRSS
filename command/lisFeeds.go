package command

import (
	"context"
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
)

func HandlerListFeeds(s *config.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: feeds")
	}
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds: %w", err)
	}
	fmt.Println("Feeds:")
	fmt.Println("---------------------------------------")
	for _, feed := range feeds {
		user, err := s.Db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't fetch user that added the feed: %w", err)
		}
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("User: %s\n", user.Name)
		fmt.Println("---------------------------------------")
	}
	return nil
}
