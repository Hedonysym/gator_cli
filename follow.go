package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Hedonysym/gator_cli/internal/database"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: follow <feedUrl>")
	}
	feedUrl := cmd.args[0]
	currentUser, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return fmt.Errorf("error fetching current user: %v", err)
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}
	if feed.ID == uuid.Nil {
		return fmt.Errorf("feed not found: %s", feedUrl)
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %v", err)
	}
	fmt.Printf("User: %s successfully followed feed: %s\n", currentUser.Name, feed.Name)
	return nil
}
