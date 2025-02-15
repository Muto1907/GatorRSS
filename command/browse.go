package command

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
)

func HandleBrowse(s *config.State, cmd Command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.Args) == 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid argument for number of posts to browse: %w", err)
		}
	}
	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.Db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("Feed: %s from: %s\n", post.Name, post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf("--- %s ----\n", post.Title)
		fmt.Printf("%s\n", post.Description)
		fmt.Printf("More at: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
