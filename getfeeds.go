package main

import (
	"context"
	"fmt"
)

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting user for feed %s: %w", feed.Name, err)
		}
		fmt.Printf("Name: %s\nURL: %s\nUser: %v\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
