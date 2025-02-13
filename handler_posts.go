package main

import (
	"context"
	"fmt"
	"html"
	"strconv"

	"github.com/mu7ammad1951/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.args) >= 1 {
		convLimit, err := strconv.Atoi(cmd.args[0])
		limit = int32(convLimit)
		if err != nil {
			fmt.Printf("error converting argument to number, defaulting to 2: use only integers > 0")
			limit = 2
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: limit,
	})
	if err != nil {
		return fmt.Errorf("error fetching posts for user %s: %w", user.Name, err)
	}

	for _, post := range posts {
		printPost(post)
	}

	return nil
}

func printPost(post database.Post) {
	fmt.Println()
	fmt.Printf("%v\n", html.UnescapeString(post.Title))
	fmt.Printf("Published on: %v\n", post.PublishedAt.Format("2006-01-02"))
	fmt.Printf("Decription:\n%v\n", html.UnescapeString(post.Description.String))
	fmt.Println()
}
