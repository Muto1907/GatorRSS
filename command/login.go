package command

import (
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
)

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	err := s.Config.SetUser(name)
	if err != nil {
		return fmt.Errorf("login for %s failed: %w", name, err)
	}
	fmt.Printf("Username has been Set to %s\n", name)
	return nil
}
