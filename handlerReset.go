
package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not delete users: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}

