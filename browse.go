package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Hedonysym/gator_cli/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) >= 2 {
		return fmt.Errorf("usage: browse <page_limit(optional)>")
	}
	var limitValue int
	var err error
	limit := 2
	if len(cmd.args) > 0 {
		limitValue, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid page limit: %v", err)
		}
		limit = limitValue
	}
	posts, err := s.db.GetPostsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching posts: %v", err)
	}
	for i, post := range posts {
		if i >= limit {
			break
		}
		fmt.Printf("%d. %s\n", i+1, post.Title)
		fmt.Printf("   URL: %s\n", post.Url)
		fmt.Printf("   Description: %v\n", post.Description)
		fmt.Printf("   Published at: %v\n", post.PublishedAt)
	}
	return nil
}
