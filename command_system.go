package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/toyeafo/blog-aggregator-go/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please provide a username")
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user %w", err)
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please provide a username")
	}

	username := cmd.Args[0]
	ctx := context.Background()

	newUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		Name:      sql.NullString{String: username},
		ID:        int32(uuid.New().ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return fmt.Errorf("error creating user %w", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting user %w", err)
	}

	fmt.Printf("User created. Details: %v", newUser)
	return nil
}
