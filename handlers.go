package main

import (
	"context"
	"errors"
	"fmt"
	"html"
	"os"
	"time"

	"github.com/Giira/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("error: login function requires a single word username")
	}
	name := cmd.arguments[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		fmt.Printf("error: user '%s' not in table\n", name)
		os.Exit(1)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error: username could not be set - %v", err)
	}
	fmt.Printf("Username set to: %s\n", s.cfg.Current_user_name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("error: register function requires a single word username")
	}
	name := cmd.arguments[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error: username could not be set - %v", err)
	}
	fmt.Printf("User created:\nID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName: %s\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	for _, user := range users {
		if user == s.cfg.Current_user_name {
			user = user + " (current)"
		}
		user = "* " + user
		fmt.Println(user)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	title := html.UnescapeString(feed.Channel.Title)
	fmt.Printf("Title: %v\n", title)
	fmt.Printf("Link: %v\n", feed.Channel.Link)
	desc := html.UnescapeString(feed.Channel.Description)
	fmt.Printf("Description: %v\n", desc)
	for i, item := range feed.Channel.Item {
		fmt.Printf("\nItem %v:\n\n", i+1)
		title = html.UnescapeString(item.Title)
		fmt.Printf("Title: %v\n", title)
		fmt.Printf("Link: %v\n", item.Link)
		desc = html.UnescapeString(item.Description)
		fmt.Printf("Description: %v\n", desc)
		fmt.Printf("Publication Date: %v\n", item.PubDate)
	}

	return nil
}
