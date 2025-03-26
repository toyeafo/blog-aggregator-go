package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/toyeafo/blog-aggregator-go/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		userID, err := s.db.GetUser(context.Background(), sql.NullString{String: s.cfg.User_name, Valid: true})
		if err != nil {
			return fmt.Errorf("error retrieving user id: %w", err)
		}
		return handler(s, c, userID)
	}
}
