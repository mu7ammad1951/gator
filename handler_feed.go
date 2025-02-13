package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mu7ammad1951/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Println(*rssFeed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("missing argument - USAGE> addFeed <name> <url>")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	params := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    currentUser.ID,
	}

	feed, err := s.db.AddFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)

	if err = handlerFeedFollow(s, command{
		name: "follow",
		args: []string{cmd.args[1]},
	}); err != nil {
		return err
	}

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	rssFeeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range rssFeeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed: %v\nURL: %v\nAdded By: %v\n\n", feed.Name, feed.Url, user.Name)
	}
	return nil
}
