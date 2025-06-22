package main

import (
	"context"
	"fmt"
	"database/sql"
	"github.com/trolioSFG/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Missing login name")
	}


	data := database.User{}
	data, err := s.db.GetUser(context.Background(), cmd.args[0])


	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("User %s not registered in db", cmd.args[0])
		}
		return err
	}


	s.cfg.SetUser(cmd.args[0])
	fmt.Println(data)
	fmt.Printf("User %s set\n", cmd.args[0])

	return nil
}

