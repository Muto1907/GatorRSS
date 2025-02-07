package command

import (
	"context"
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
)

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	user, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldnt find user: %w", err)
	}
	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("setting user %s failed: %w", name, err)
	}
	fmt.Printf("Username has been Set to %s\n", name)
	return nil
}
