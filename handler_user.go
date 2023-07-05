package main

import (
	"encoding/json"
	"fmt"
	"kimb/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

/**REGISTER USER */
func (apiCfg *apiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the values: %s", err))
		return
	}

	// PASWORDGEN
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error with password encryption: %s", err))
		return
	}

	// SAVE USER DATA
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:          uuid.New(),
		Name:        params.Name,
		DisplayName: params.DisplayName,
		Email:       params.Email,
		Password:    hash,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error during the registration of the user: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, returnUser(user))
}

/**GET USER FROM API KEY */
func (apiCfg *apiConfig) handlerGetUserFromAPI(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, returnUser(user))
}

/**GET USER DATA FROM NAME */
func (apiCfg *apiConfig) handlerGetUserFromName(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	userInfo, err := apiCfg.DB.GetUserInfo(r.Context(), username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting the user: %s", err))
		return
	}
	respondWithJSON(w, http.StatusAccepted, userInfo)
}
