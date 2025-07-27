package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Hedonysym/gator_cli/internal/database"
)

func scrapeFeeds(s *state, user database.User) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}
	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            feed.ID,
	}
	feed, err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}
	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	if len(rss.Channel.Item) == 0 {
		fmt.Printf("No items found in feed %s\n", feed.Url)
		return nil
	}
	fmt.Printf("Fetched %d items from feed %s\n", len(rss.Channel.Item), feed.Url)
	for _, item := range rss.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}
