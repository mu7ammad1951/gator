package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mu7ammad1951/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("missing argument - USAGE: agg <duration-string>")
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid <duration-string> : use <X>s or <X>m or <X>h\n<X> is an integer")
	}

	fmt.Printf("Collecting feeds every %v", duration.String())
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("error scraping feeds: %w", err)
		}
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("missing arguments - USAGE: addFeed <name> <url>")
	}

	params := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.AddFeed(context.Background(), params)
	if err != nil {
		return err
	}

	dataRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Println("successfully following!")
	fmt.Printf("User: %v\nFeed: %v", dataRow.UserName, dataRow.FeedName)

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
