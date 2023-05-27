package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nguyenanhhao221/go-rss/internal/database"
)

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, error := apiCfg.DB.GetFeeds(r.Context())
	if error != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot get feeds: %v", error))
		return
	} else {
		convertedFeeds := databaseFeedsToFeeds(feeds)
		responseWithJSON(w, 200, convertedFeeds)
	}
}

// handlerFeed a method and pass it to the pointer of the apiConfig struct.
// This way it have access to our database
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot create feed: %v", err))
		return
	} else {
		convertedFeed := databaseFeedToFeed(feed)
		responseWithJSON(w, 201, convertedFeed)
	}
}
