package command

import (
	"context"
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
)

func HandleUnfollow(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <feed_url>")
	}
	feed, err := s.Db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't fetch feed to unfollow: %w", err)
	}
	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.Db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't unfollow: %w", err)
	}
	fmt.Printf("%s unfollowed successfully\n", feed.Name)
	return nil
}
