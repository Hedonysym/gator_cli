package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: following")
	}
	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
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
