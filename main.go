package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Giira/blogaggregator/internal/config"
	"github.com/Giira/blogaggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	s := state{}
	s.cfg = &cfg

	db, err := sql.Open("postgres", s.cfg.Db_url)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	s.db = dbQueries

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)

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
