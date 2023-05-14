package main

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestPortEnv(t *testing.T) {

	// Load the env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading env")
	}

	port := os.Getenv("PORT")

	// Check if the PORT environment variable was correctly retrieved
	if port != "8000" {
		t.Errorf("PORT environment variable not retrieved correctly, got: %s, want: %s", os.Getenv("PORT"), "8000")
	}
}
