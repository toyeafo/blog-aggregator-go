package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/toyeafo/blog-aggregator-go/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	userID, err := s.db.GetUser(context.Background(), sql.NullString{String: s.cfg.User_name, Valid: true})
	if err != nil {
		return fmt.Errorf("error retrieving user id: %w", err)
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("please provide the name and url of the field")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      sql.NullString{String: cmd.Args[0], Valid: true},
		Url:       cmd.Args[1],
		UserID:    userID.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error creating feed %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    userID.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow record and retrieving field: %w", err)
	}

	fmt.Printf("Feed and follow records created. Details: %v\n%v", feed, feedFollow)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithName(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving feeds with usernames: %w", err)
	}

	for i := range feeds {
		fmt.Printf("* ID:            %s\n", feeds[i].ID)
		fmt.Printf("* Created:       %v\n", feeds[i].CreatedAt)
		fmt.Printf("* Updated:       %v\n", feeds[i].UpdatedAt)
		fmt.Printf("* Name:          %s\n", feeds[i].Name.String)
		fmt.Printf("* URL:           %s\n", feeds[i].Url)
		fmt.Printf("* User:          %s\n", feeds[i].Name_2.String)
	}
	return nil
}
