package main

import (
	"fmt"
	"os"

	"github.com/Hedonysym/gator_cli/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("No command provided")
		os.Exit(1)
	}
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	state := &state{config: &cfg}
	cmds := commands{}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(state, cmd)
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
