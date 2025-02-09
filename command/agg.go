package command

import (
	"context"
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/rss"
)

func HandleAgg(s *config.State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch RSSFeed: %w", err)
	}
	fmt.Println(feed)
	return nil
}
