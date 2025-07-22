package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Hedonysym/gator_cli/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username argument is required")
	}
	username := cmd.args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("User registered successfully")
	fmt.Println(user)
	return nil
}
