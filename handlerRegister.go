
package main

import (
	"fmt"
	"time"
	"github.com/trolioSFG/database"
	"context"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Missing register name")
	}

	// New user in DB
	data := database.CreateUserParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), data)
	if err != nil {
		return err
	}

	s.cfg.SetUser(cmd.args[0])
	fmt.Printf("User %s was created\n", cmd.args[0])
	fmt.Printf("%+v\n", user)


	return nil
}

