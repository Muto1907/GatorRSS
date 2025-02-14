package command

import (
	"context"
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
)

func MiddleWareLoggedIn(handler func(s *config.State, cmd Command, user database.User) error) func(*config.State, Command) error {

	return func(s *config.State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUsername)
		if err != nil {
			return fmt.Errorf("couldn't fetch user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
