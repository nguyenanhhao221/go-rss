package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	// Allow cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Add router handler
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	// mount the v1Router to the /v1 route
	// so if we access /v1/healthz the handlerReadiness will be called
	router.Mount("/v1", v1Router)

	// Start the server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server is starting on port %s", port)
	serverErr := srv.ListenAndServe()
	if serverErr != nil {
		log.Fatal("Error while Listen and Serve the server")
	}
}
