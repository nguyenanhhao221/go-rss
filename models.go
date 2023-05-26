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
