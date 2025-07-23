package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Hedonysym/gator_cli/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}
	name := cmd.args[0]
	url := cmd.args[1]

	params := database.AddFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feedID, err := s.db.AddFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error adding feed: %w", err)
	}

	fmt.Printf("Feed added with ID: %s\n", feedID)
	return nil
}
