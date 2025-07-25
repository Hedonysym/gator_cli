package main

import (
	"context"
	"fmt"

	"github.com/Hedonysym/gator_cli/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: unfollow <feed_url>")
	}
	feedUrl := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error retrieving feed: %v", err)
	}
	params := database.FeedUnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.FeedUnfollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}
	fmt.Printf("Successfully unfollowed feed: %s\n", feedUrl)
	return nil

}
