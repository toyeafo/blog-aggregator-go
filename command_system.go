package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/toyeafo/blog-aggregator-go/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please provide a username")
	}

	username := cmd.Args[0]
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, sql.NullString{String: username, Valid: true})
	if err == sql.ErrNoRows {
		fmt.Println("User doesn't exist in the database: ", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(username)
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

	_, err := s.db.GetUser(ctx, sql.NullString{String: username, Valid: true})
	if err == nil {
		os.Exit(1)
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("error checking user existence: %w", err)
	}

	newUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		Name:      sql.NullString{String: username, Valid: true},
		ID:        uuid.New(),
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
