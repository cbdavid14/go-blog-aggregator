package main

import (
	"database/sql"
	"github.com/cbdavid14/go-blog-aggregator/internal/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := mux.NewRouter()

	mux.HandleFunc("/v1/user", apiCfg.handlerUserCreate).Methods("POST")
	mux.HandleFunc("/v1/user", apiCfg.middlewareAuth(apiCfg.handlerUserGet)).Methods("GET")
	mux.HandleFunc("/v1/users", apiCfg.handlerUserGetAll).Methods("GET")

	mux.HandleFunc("/v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate)).Methods("POST")
	mux.HandleFunc("/v1/feeds", apiCfg.handleGetFeeds).Methods("GET")

	mux.HandleFunc("/v1/feed-follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete)).Methods("DELETE")
	mux.HandleFunc("/v1/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate)).Methods("POST")
	mux.HandleFunc("/v1/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGet)).Methods("GET")

	mux.HandleFunc("/v1/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsGet)).Methods("GET")
	mux.HandleFunc("/v1/posts/all", apiCfg.handlerPostsGetAll).Methods("GET")
	mux.HandleFunc("/v1/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsCreate)).Methods("POST")

	mux.HandleFunc("/v1/healthz", HandlerReadiness).Methods("GET")
	mux.HandleFunc("/v1/err", HandlerErr).Methods("GET")

	corsMux := CorsMiddleware(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	//go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
