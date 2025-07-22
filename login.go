package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username argument is required")
	}
	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		defer os.Exit(1)
		return fmt.Errorf("user not found: %w", err)
	}
	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Logged in as %s\n", username)
	return nil
}
