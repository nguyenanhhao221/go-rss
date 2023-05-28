package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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

	feed_follows, err := apiCfg.DB.CreateFeedsFollow(r.Context(), database.CreateFeedsFollowParams{
		ID:        uuid.New(),
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot create feed follow: %v", err))
		return
	} else {
		convertedFeedFollow := databaseFeedFollowToFeedFollow(feed_follows)
		responseWithJSON(w, 201, convertedFeedFollow)
	}
}

func (apiCfg *apiConfig) handlerGetFeedToFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiCfg.DB.GetFeedToFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot create feed follow: %v", err))
		return
	} else {
		convertedFeedFollows := databaseFeedFollowsToFeedFollows(feed_follows)
		responseWithJSON(w, 200, convertedFeedFollows)
	}
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error while parsing uuid: %v", err))
		return
	}

	deleteFeedFollowErr := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{ID: feedFollowID, UserID: user.ID})

	if deleteFeedFollowErr != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot delete feed follow: %v", deleteFeedFollowErr))
		return
	}

	responseWithJSON(w, 204, struct{}{})
}
