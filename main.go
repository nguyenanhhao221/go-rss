package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nguyenanhhao221/go-rss/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// The URL in .env to connect to SQL database
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// Open a connection to the postgres database
	// sql.Open is used to establish a connection to the PostgreSQL database.
	// However, the sql.Open function only creates a connection object, it doesn't actually establish a connection to the database.
	sqlConnection, dbConnErr := sql.Open("postgres", dbURL)
	if dbConnErr != nil {
		log.Fatal("Cannot connect to the database: ", dbConnErr)
	}
	// sql.Open successfully returns an instance of sql.DB regardless of whether the database server is running or not.
	// To check if the connection was successful, you need to call the Ping method on the sql.DB instance.
	if err := sqlConnection.Ping(); err != nil {
		log.Fatal("Failed to ping the database, did you forget to run Docker? Error: ", err)
	}
	// database is the package generate by sqlc which contain our queries to the actual database.
	// So basically here The queries object is responsible for executing SQL queries against the database using the underlying connection.
	queries := database.New(sqlConnection)

	// This will get pass to our handler so that the handler have access to the database
	apiCfg := apiConfig{
		DB: queries,
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
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	// Call the middlewareAuth function and pass the handler to that function, similar concept to callback, higher order function
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))
	// feeds
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	// Feed follow
	v1Router.Post("/feeds_follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedsFollow))
	v1Router.Get("/feeds_follow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedToFollow))
	v1Router.Delete("/feeds_follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
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
