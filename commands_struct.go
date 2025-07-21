package main

import "github.com/Hedonysym/gator_cli/internal/config"

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}
