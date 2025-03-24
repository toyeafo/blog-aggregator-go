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

	err = s.cfg.SetUser(newUser.Name.String)
	if err != nil {
		return fmt.Errorf("error setting user %w", err)
	}

	fmt.Printf("User created. Details: %v", newUser)
	return nil
}

func resetHandler(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users from the table: %w", err)
	}
	return nil
}

func users(s *state, cmd command) error {
	usernames, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retueving users from database: %w", err)
	}

	for i := 0; i < len(usernames); i++ {
		if usernames[i].Name.String == s.cfg.User_name {
			fmt.Println(usernames[i].Name.String, "(current)")
		} else {
			fmt.Println(usernames[i].Name.String)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	aggFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}
	fmt.Printf("feed has been fetched: %+v\n", aggFeed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	user := s.cfg.User_name
	userID, err := s.db.GetUser(context.Background(), sql.NullString{String: user, Valid: true})
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

	fmt.Printf("Feed created. Details: %v", feed)
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

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please provide the url to follow")
	}

	url := cmd.Args[0]
	user := s.cfg.User_name
	userID, err := s.db.GetUser(context.Background(), sql.NullString{String: user, Valid: true})
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

func handlerFollowing(s *state, cmd command) error {
	user := s.cfg.User_name
	userID, err := s.db.GetUser(context.Background(), sql.NullString{String: user, Valid: true})
	if err != nil {
		return fmt.Errorf("error retrieving user id: %w", err)
	}

	following, err := s.db.GetFeedsFollowForUser(context.Background(), sql.NullString{String: userID.Name.String, Valid: true})
	if err != nil {
		return fmt.Errorf("error retrieving feeds of user %v, problem: %v", user, err)
	}

	for f := range following {
		fmt.Println(following[f].FeedName.String)
	}
	return nil
}
