package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/toyeafo/blog-aggregator-go/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please provide the url to follow")
	}

	url := cmd.Args[0]
	userID, err := s.db.GetUser(context.Background(), sql.NullString{String: s.cfg.User_name, Valid: true})
	if err != nil {
		return fmt.Errorf("error retrieving user id: %w", err)
	}

	feedID, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error retrieving feed id of the url: %w", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    userID.ID,
		FeedID:    feedID.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow record and retrieving field: %w", err)
	}

	fmt.Printf("Name of field: %s", follow.FeedName.String)
	fmt.Printf("Current User: %s", follow.UserName.String)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	userID, err := s.db.GetUser(context.Background(), sql.NullString{String: s.cfg.User_name, Valid: true})
	if err != nil {
		return fmt.Errorf("error retrieving user id: %w", err)
	}

	following, err := s.db.GetFeedsFollowForUser(context.Background(), userID.ID)
	if err != nil {
		return fmt.Errorf("error retrieving feeds of user %v, problem: %v", s.cfg.User_name, err)
	}

	for f := range following {
		fmt.Println(following[f].FeedName.String)
	}
	return nil
}
