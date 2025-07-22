package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}
	fmt.Println("All users have been reset successfully")
	return nil
}
