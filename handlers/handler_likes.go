package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

/**LIKE POST */
func (apiCfg *apiConfig) handlerLikePost(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		PostID uuid.UUID `json:"post_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the values: %s", err))
		return
	}

	// SAVE USER DATA
	likedPost, err := apiCfg.DB.LikePost(r.Context(), database.LikePostParams{
		ID:     uuid.New(),
		UserID: user.ID,
		PostID: params.PostID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error liking the post: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, returnLikedPost(likedPost))
}

/**GET ALL OF USER LIKES */
func (apiCfg *apiConfig) handlerGetAllUserLikes(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the values: %s", err))
		return
	}

	// Get all the posts the user liked
	allLikedPosts, err := apiCfg.DB.GetLikedPosts(r.Context(), params.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting user likes: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, returnAllLikedPosts(allLikedPosts))
}

/**GET POST LIKES */
func (apiCfg *apiConfig) handlerGetPostLikes(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		PostID uuid.UUID `json:"post_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the values: %s", err))
		return
	}

	// Get all the posts the user liked
	allPostLikes, err := apiCfg.DB.GetLikesOfPost(r.Context(), params.PostID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting post likes: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, returnmPostLikes(allPostLikes))
}

/**UNLIKE A POST */
func (apiCfg *apiConfig) handlerRemovePostLike(w http.ResponseWriter, r *http.Request, user database.User) {
	postIDstr := chi.URLParam(r, "postID")

	postID, err := uuid.Parse(postIDstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the post id: %s", err))
		return
	}

	err = apiCfg.DB.UnlikePost(r.Context(), database.UnlikePostParams{
		PostID: postID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error unliking the post: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
