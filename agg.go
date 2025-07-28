package main

import (
	"fmt"
	"time"

	"github.com/Hedonysym/gator_cli/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: agg <*m*s>(minutes/seconds)")
	}
	timeBetweenFetches, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", cmd.args[0])
	ticker := time.NewTicker(timeBetweenFetches)
	for ; ; <-ticker.C {
		scrapeFeeds(s, user)
	}
	return nil
}
