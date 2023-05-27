package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/nguyenanhhao221/go-rss/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

// databaseUserToUser convert the type User originally from database to our own User type that better match the json type
// for example "create_at" instead of CreatedAt
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreateAt:  dbUser.CreateAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreateAt:  dbFeed.CreateAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedsFollow struct {
	ID        uuid.UUID `json:"id"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

// databaseFeedFollowToFeedFollow convert type to json case
func databaseFeedFollowToFeedFollow(dbFeedsFollow database.FeedsFollow) FeedsFollow {
	return FeedsFollow{
		ID:        dbFeedsFollow.ID,
		CreateAt:  dbFeedsFollow.CreateAt,
		UpdatedAt: dbFeedsFollow.UpdatedAt,
		UserID:    dbFeedsFollow.UserID,
		FeedID:    dbFeedsFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedsFollows []database.FeedsFollow) []FeedsFollow {
	feedToFollows := []FeedsFollow{}
	for _, feedToFollow := range dbFeedsFollows {
		feedToFollows = append(feedToFollows, databaseFeedFollowToFeedFollow(feedToFollow))
	}
	return feedToFollows
}
