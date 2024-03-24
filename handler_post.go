package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cbdavid14/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}
	fmt.Printf("handlerPostsGet User: %s\n", user.ID)
	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get posts [%s]", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}

func (cfg *apiConfig) handlerPostsGetAll(writer http.ResponseWriter, request *http.Request) {
	posts, err := cfg.DB.GetAllPost(request.Context())
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, fmt.Sprintf("Couldn't get posts [%s]", err))
		return
	}

	respondWithJSON(writer, http.StatusOK, databasePostsToPosts(posts))

}

func (cfg *apiConfig) handlerPostsCreate(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		Title       string `json:"title"`
		Url         string `json:"url"`
		Description string `json:"description"`
		FeedID      string `json:"feed_id"`
	}
	decoder := json.NewDecoder(request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, fmt.Sprintf("Couldn't decode parameters [%s]", err))
		return
	}
	resul, err := cfg.DB.CreatePost(request.Context(), database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Title:       params.Title,
		Url:         params.Url,
		Description: sql.NullString{String: params.Description, Valid: true},
		PublishedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		FeedID:      uuid.MustParse(params.FeedID),
	})
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, fmt.Sprintf("Couldn't create post [%s]", err))
		return
	}
	respondWithJSON(writer, http.StatusOK, databasePostToPost(resul))
}
