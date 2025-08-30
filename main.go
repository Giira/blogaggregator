package main

import (
	"fmt"
	"os"

	"github.com/Giira/blogaggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	s := state{}
	s.cfg = &cfg

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("error: no function called\n")
		os.Exit(1)
	} else if len(args) == 2 && args[1] == "login" {
		fmt.Printf("error: login requires a username\n")
		os.Exit(1)
	}
	cmd := command{
		name:      args[1],
		arguments: args[2:],
	}

	err = commands.run(&s, cmd)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
