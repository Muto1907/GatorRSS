package command

import (
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*config.State, Command) error
}

func (c *Commands) Register(Name string, f func(*config.State, Command) error) {
	c.Handlers[Name] = f
}

func (c *Commands) Run(s *config.State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("no such Command as %s found", cmd.Name)
	}
	return handler(s, cmd)

}
