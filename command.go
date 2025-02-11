package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(s *state, cmd command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.handlers[cmd.name]
	if !exists {
		return fmt.Errorf("invalid command")
	}
	return f(s, cmd)
}
