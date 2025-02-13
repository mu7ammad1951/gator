package main

import (
	"context"
	"fmt"

	"github.com/mu7ammad1951/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {

	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error fetching current user: %w", err)
		}

		return handler(s, cmd, currentUser)

	}
}
