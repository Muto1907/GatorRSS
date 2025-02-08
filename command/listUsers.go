package command

import (
	"context"
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
)

func HandlerListUsers(s *config.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: users")
	}
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get users: %w", err)
	}
	for _, user := range users {
		if s.Config.CurrentUsername == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
