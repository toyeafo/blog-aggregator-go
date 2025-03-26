package main

import (
	"context"
	"fmt"

	"github.com/toyeafo/blog-aggregator-go/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	aggFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}
	fmt.Printf("feed has been fetched: %+v\n", aggFeed)
	return nil
}

func scrapeFeeds(s *state, cmd command) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: nextFeed.LastFetchedAt,
		UpdatedAt:     nextFeed.UpdatedAt,
		ID:            nextFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("error marking feed as fetched in database: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error retrieving feed by URL: %w", err)
	}

	for i := range feed.Name.String {
		fmt.Println(feed.Name.String[i])
	}
	return nil
}
