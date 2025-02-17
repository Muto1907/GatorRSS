package command

import (
	"context"
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
)

func HandlerFollowing(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: following")
	}
	followList, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't fetch follow list: %w", err)
	}
	fmt.Printf("Followed Feeds:\n")
	for _, entry := range followList {
		fmt.Printf("Feed: %s Followed By: %s\n", entry.FeedName, entry.UsersName)
	}
	return nil
}
