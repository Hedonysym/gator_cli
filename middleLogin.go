package main

import (
	"context"
	"fmt"

	"github.com/Hedonysym/gator_cli/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		current, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
		if err != nil {
			return fmt.Errorf("error retrieving current user: %v", err)
		}
		return handler(s, cmd, current)
	}

}
