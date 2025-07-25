package main

import (
	"context"
	"fmt"

	"github.com/Hedonysym/gator_cli/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: following")
	}
	ffs, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching followings: %w", err)
	}

	fmt.Printf("%s's feeds:\n", user.Name)
	for _, f := range ffs {
		fmt.Printf("* %s\n", f.FeedName)
	}
	return nil
}
