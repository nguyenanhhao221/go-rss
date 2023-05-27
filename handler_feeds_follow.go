package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nguyenanhhao221/go-rss/internal/database"
)

// handlerCreateFeedsFollow create a feed follow base on feed_id(from client) and user_id(auth with api key)
func (apiCfg *apiConfig) handlerCreateFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed_follows, error := apiCfg.DB.CreateFeedsFollow(r.Context(), database.CreateFeedsFollowParams{
		ID:        uuid.New(),
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if error != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot create feed follow: %v", err))
		return
	} else {
		convertedFeedFollow := databaseFeedFollowToFeedFollow(feed_follows)
		responseWithJSON(w, 201, convertedFeedFollow)
	}
}
