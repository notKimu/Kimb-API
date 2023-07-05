package main

import (
	"encoding/json"
	"fmt"
	"kimb/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

/**LIKE POST */
func (apiCfg *apiConfig) handlerFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		UserID     uuid.UUID `json:"user_id"`
		FollowedID uuid.UUID `json:"followed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the values: %s", err))
		return
	}

	if user.ID == params.FollowedID {
		respondWithError(w, http.StatusBadRequest, "I wont comment on this...")
		return
	}

	// SAVE FOLLOW
	followedUser, err := apiCfg.DB.FollowUser(r.Context(), database.FollowUserParams{
		ID:         uuid.New(),
		UserID:     user.ID,
		FollowedID: params.FollowedID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error liking the post: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, returnFollow(followedUser))
}

/**GET ALL OF USER LIKES */
func (apiCfg *apiConfig) handlerGetFollowing(w http.ResponseWriter, r *http.Request) {
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
	following, err := apiCfg.DB.GetFollowing(r.Context(), params.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting user following: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, returnFollowing(following))
}

/**GET POST LIKES */
func (apiCfg *apiConfig) handlerGetFollowers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FollowedID uuid.UUID `json:"followed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the values: %s", err))
		return
	}

	// Get all the posts the user liked
	followers, err := apiCfg.DB.GetFollowers(r.Context(), params.FollowedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting post likes: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, returnFollowers(followers))
}

/**UNLIKE A POST */
func (apiCfg *apiConfig) handlerUnfollow(w http.ResponseWriter, r *http.Request, user database.User) {
	followedIDstr := chi.URLParam(r, "followedID")

	followedID, err := uuid.Parse(followedIDstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the user id: %s", err))
		return
	}

	err = apiCfg.DB.Unfollow(r.Context(), database.UnfollowParams{
		FollowedID: followedID,
		UserID:     user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error unfollowing the user: %s", err))
		return
	}

	respondWithJSON(w, http.StatusAccepted, struct{}{})
}
