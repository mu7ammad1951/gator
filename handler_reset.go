package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetTable(context.Background())
	if err != nil {
		return fmt.Errorf("unsuccessful reset: %v", err)
	}
	fmt.Println("successfully reset the database")
	return nil
}
