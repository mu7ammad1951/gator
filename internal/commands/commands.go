package commands

import (
	"fmt"

	"github.com/mu7ammad1951/gator/internal/config"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(s *state, cmd command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	s.cfg.CurrentUserName = cmd.args[0]
	err := config.Write(*s.cfg)
	if err != nil {
		return err
	}

	fmt.Println("The username has been set")
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.handlers[cmd.name]
	if !exists {
		return fmt.Errorf("invalid command")
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
