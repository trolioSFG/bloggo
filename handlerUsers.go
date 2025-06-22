package main

import (
	"context"
	"fmt"
//	"database/sql"
//	"github.com/trolioSFG/database"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("Could not list users")
	}

	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUser {
			fmt.Printf(" (current)")
		}
		fmt.Printf("\n")
	}

	return nil
}

