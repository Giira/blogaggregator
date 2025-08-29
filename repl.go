package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Giira/blogaggregator/internal/config"
)

type state struct {
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
	// run command
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	// register handler function
}

func cleanInput(text string) []string {
	var output []string
	text = strings.ToLower(text)
	output = strings.Fields(text)
	return output
}

func handlerLogin(s *state, cmd command) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Gator > ")
		boo := scanner.Scan()

		if !boo {
			return errors.New("Error: Input required")
		} else {
			input := scanner.Text()
			i_slice := cleanInput(input)

			switch i_slice[0] {
			case "login":

			}
		}
	}
}
