package main

import (
	"fmt"
	"net/http"

	"github.com/nguyenanhhao221/go-rss/internal/auth"
	"github.com/nguyenanhhao221/go-rss/internal/database"
)

type authedHeader func(http.ResponseWriter, *http.Request, database.User)

// middlewareAuth get the user with api key
func (apiCfg *apiConfig) middlewareAuth(handler authedHeader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		//NOTE: Understand why we need to use context here
		user, userErr := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if userErr != nil {
			responseWithError(w, 400, fmt.Sprintf("Couldn't  get user: %v", userErr))
			return
		}

		handler(w, r, user)
	}
}
