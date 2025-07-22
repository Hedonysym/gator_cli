package main

import (
	"context"
	"fmt"
)

func handlerAgg(state *state, cmd command) error {
	/*
		if len(cmd.args) < 1 {
			return fmt.Errorf("usage: agg <feed_url>")
		}*/
	feedURL := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()
	feed, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
