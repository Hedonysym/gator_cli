package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Hedonysym/gator_cli/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
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
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID[0].ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("Feed added")
	return nil
}
