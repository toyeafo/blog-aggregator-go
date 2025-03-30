package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/toyeafo/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please provide the url to follow")
	}

	url := cmd.Args[0]
	feedID, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error retrieving feed id of the url: %w", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
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
	following, err := s.db.GetFeedsFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error retrieving feeds of user %v, problem: %v", s.cfg.User_name, err)
	}

	for f := range following {
		fmt.Println(following[f].FeedName.String)
	}
	return nil
}

func handlerDeleteFollow(s *state, cmd command, user database.User) error {
	url := cmd.Args[0]
	feedID, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error retrieving feed id of the url: %w", err)
	}
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedID.ID,
	})
	if err != nil {
		return fmt.Errorf("issue deleting feed follow from database: %w", err)
	}
	return nil
}
