package main

import (
	"fmt"
)

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.commands[cmd.name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.commands == nil {
		c.commands = make(map[string]func(*state, command) error)
	}
	c.commands[name] = f
}
