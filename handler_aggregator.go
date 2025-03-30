package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/toyeafo/blog-aggregator-go/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error converting argument to duration: %w", err)
	}
	log.Printf("Collecting feeds every %s\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			scrapeFeeds(s)
		case <-context.Background().Done():
			return nil
		}

	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving next feed: %w", err)
	}

	aggFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched in database: %w", err)
	}

	for _, item := range aggFeed.Channel.Item {
		t, _ := time.Parse(time.RFC1123Z, item.PubDate)

		existingPost, err := s.db.GetPostByUrl(context.Background(), item.Link)
		if err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("error checking for existing post: %w", err)
		}

		if existingPost != nil {
			continue
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: t, Valid: true},
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			return fmt.Errorf("error adding post to database: %w", err)
		}
	}
	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	limitVal := 2
	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limitVal = specifiedLimit
		} else {
			return fmt.Errorf("please provide a valid number: %w", err)
		}
	}
	post, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limitVal)})
	if err != nil {
		return fmt.Errorf("error retrieving posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(post), user.Name.String)
	for _, item := range post {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Description: %s\n", item.Description.String)
		fmt.Printf("Published At: %v\n", item.PublishedAt.Time.Format("Mon Jan 2"))
	}
	return nil
}
