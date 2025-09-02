package main

import (
	"errors"
	"fmt"

	"github.com/Giira/blogaggregator/internal/config"
	"github.com/Giira/blogaggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commands[cmd.name]
	if !ok {
		return errors.New("error: no such function")
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("error: login function requires a single word username")
	}
	name := cmd.arguments[0]
	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error: username could not be set - %v", err)
	}
	fmt.Printf("Username set to: %s\n", s.cfg.Current_user_name)
	return nil
}
