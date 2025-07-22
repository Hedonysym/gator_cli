package main

import (
	"github.com/Hedonysym/gator_cli/internal/config"
	"github.com/Hedonysym/gator_cli/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}
