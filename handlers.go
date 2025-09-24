package main

import (
	"context"
	"errors"
	"fmt"
	"log"
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

	return nil
}

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return errors.New("error: addfeed requires a feed name and a url")
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Printf("Feed created:\nID: %v\nCreated at: %v\nUpdated at: %v\n", feed.ID, feed.CreatedAt, feed.UpdatedAt)
	fmt.Printf("Feed name: %v\nUrl: %v\nUserID: %v\n", feed.Name, feed.Url, feed.UserID)

	cmdff := command{
		name:      "follow",
		arguments: []string{url},
	}
	err = handlerFollow(s, cmdff, user)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	for i, feed := range feeds {
		fmt.Printf("Feed %v\nName: %v\nUrl: %v\nUser: %v\n", i+1, feed.Name, feed.Url, feed.Name_2)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("error: follow function requires a single url")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	fmt.Printf("Feed followed:\nFeed: %v\nUser: %v\n", feedFollow.FeedName, feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return errors.New("error: following command takes no arguments")
	}
	fffu, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	for _, user_follows := range fffu {
		fmt.Printf("* %v\n", user_follows.FeedName)
	}
	return nil
}

func midLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		return handler(s, cmd, user)
	}
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("error: unfollow command takes a single url")
	}
	err := s.db.Unfollow(context.Background(), database.UnfollowParams{
		Url:  cmd.arguments[0],
		Name: user.Name,
	})
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
