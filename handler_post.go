package main

import (
	"encoding/json"
	"fmt"
	"kimb/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

/**CREATE POST */
func (apiCfg *apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Content string `json:"content"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the values: %s", err))
		return
	}

	// SAVE USER DATA
	post, err := apiCfg.DB.CreatePost(r.Context(), database.CreatePostParams{
		ID:      uuid.New(),
		Content: params.Content,
		UserID:  user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error posting the post u.u: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, returnPost(post))
}

/**UNLIKE A POST */
func (apiCfg *apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request, user database.User) {
	postIDstr := chi.URLParam(r, "postID")

	postID, err := uuid.Parse(postIDstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the post id: %s", err))
		return
	}

	err = apiCfg.DB.DeletePost(r.Context(), database.DeletePostParams{
		ID:     postID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting the post: %s", err))
		return
	}

	respondWithJSON(w, http.StatusAccepted, struct{}{})
}
