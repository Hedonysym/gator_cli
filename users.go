package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	if len(users) == 0 {
		fmt.Println("No users found")
		return nil
	}
	fmt.Println("Registered users:")
	for _, user := range users {
		if user.Name == s.config.Current_user_name {
			fmt.Println("* ", user.Name, "(current)")
			continue
		} else {
			fmt.Println("* ", user.Name)
		}
	}
	return nil
}
