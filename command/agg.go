package command

import (
	"fmt"
	"time"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/rss"
)

func HandleAgg(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}
	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse time between reqs: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", duration.String())
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		err = rss.ScrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("couldn't scrape feed: %w", err)
		}
	}

}
