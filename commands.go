package main

import (
	"fmt"
)


type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.cmds[cmd.name]; !ok {
		return fmt.Errorf("Unknown command %s", cmd.name)
	}

	return c.cmds[cmd.name](s, cmd)
}

func (c *commands) register(name string, f func(*state, command)error) {
	c.cmds[name] = f
}

