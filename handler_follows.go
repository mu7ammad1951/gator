package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mu7ammad1951/gator/internal/database"
)

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("missing arguments - USAGE: follow <url>")
	}
	url := cmd.args[0]

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed does not exist - %w, add using 'addfeed <url>'", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}
	dataRows, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}
	dataRow := dataRows[0]
	fmt.Printf("%v is now following %v\n", dataRow.UserName, dataRow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching current user information: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error fetching feeds that user %v follows: %w", currentUser.Name, err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", currentUser.Name)
	for _, feed := range feedFollows {
		fmt.Printf("%v\n", feed.FeedName)
	}

	return nil
}
