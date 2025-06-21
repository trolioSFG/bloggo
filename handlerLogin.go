package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Missing login name")
	}

	s.cfg.SetUser(cmd.args[0])
	fmt.Printf("User %s set\n", cmd.args[0])

	return nil
}

