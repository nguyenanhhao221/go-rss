package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nguyenanhhao221/go-rss/internal/database"
	"github.com/nguyenanhhao221/go-rss/util"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scrapping on %v goroutines on every %s duration", concurrency, timeBetweenRequest)
	// Similar to interval
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error while fetching feeds: ", err)
			// Don't stop even with error because we need this function to keep running
			continue
		}

		// Createa a wait group to use for go routines
		wg := &sync.WaitGroup{}
		// Loop over the feeds
		// For each feed, we add 1 to the wait group
		// Then we run the scrapeFeed function with take the wg, and defer the call Done
		// For example if we have 30 feeds, each feed get loop through and we add total 30 to the wg.
		// Also we spawn 30 go routines of scrapeFeed, this function will run and defer the call to wg.Done.
		// And once all 30 go routines is done, 30 wg.Done is called and make the wait group stop
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		// Block until the wait group counter is 0
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	// because of defer, this will always be call last in this function
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error while marking feed as fetched", err)
	}

	rssFeed, urlFetchErr := urlToFeed(feed.Url)

	if urlFetchErr != nil {
		log.Println("Error while fetching for xml", urlFetchErr)
	}

	for _, item := range rssFeed.Channel.Item {

		// Convert type
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		// Parse date
		pubDate, err := util.ParsePubDate(item.PubDate)
		if err != nil {
			log.Printf("Could not parse date %v with err %v", item.PubDate, err)
		}

		_, createPostErr := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreateAt:    time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubDate,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if createPostErr != nil {
			if strings.Contains(createPostErr.Error(), "duplicate key") {
				continue
			}
			log.Printf("Could not create post %v", createPostErr)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
