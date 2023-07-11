package main

import (
	"encoding/json"
	"fmt"
	"kimb/internal/database"
	"net/http"
	"net/mail"
	"time"

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

	// Check if email is valid
	_, err = mail.ParseAddress(params.Email)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid email adress")
		return
	}
	// Check the length of the names
	if len(params.Name) < 3 || len(params.DisplayName) < 3 || len(params.Name) > 20 || len(params.DisplayName) > 20 {
		respondWithError(w, http.StatusForbidden, "Names must be between 3 and 20 characters long >:c")
		return
	}
	// Check password security
	if len(params.Password) < 8 || len(params.Password) > 50 {
		respondWithError(w, http.StatusForbidden, "Your password must be between 8 and 50 characters!")
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

	// Get the user.apiKey and store it in the cookies as Authorization
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "ApiKey " + user.ApiKey,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Change for prod XD
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)

	respondWithJSON(w, http.StatusCreated, returnUser(user))
}

/**LOGIN USER */
func (apiCfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing the values: %s", err))
		return
	}

	// COMPARE USER DATA
	userData, err := apiCfg.DB.GetUserFromLogin(r.Context(), params.Identifier)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error logging in: %s", err))
		return
	}

	comparePasswordsErr := bcrypt.CompareHashAndPassword(userData.Password, []byte(params.Password))
	if comparePasswordsErr != nil {
		respondWithError(w, http.StatusForbidden, "Incorrect login")
		return
	}

	// Get the user.apiKey and store it in the cookies as Authorization
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "ApiKey " + userData.ApiKey,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Change for prod XD
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)

	respondWithJSON(w, http.StatusAccepted, "Logged in! Enjoy -w-")
}

/**GET USER FROM API KEY */
func (apiCfg *apiConfig) handlerGetUserFromAPI(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, returnUser(user))
}

/**GET USER DATA FROM NAME OR ID */
func (apiCfg *apiConfig) handlerGetUserFromNameOrID(w http.ResponseWriter, r *http.Request) {
	type UserPublicInfo struct {
		ID           uuid.UUID `json:"id"`
		CreatedAt    time.Time `json:"created_at"`
		Name         string    `json:"name"`
		DisplayName  string    `json:"display_name"`
		Description  string    `json:"description"`
		Verified     bool      `json:"verified"`
		Organization bool      `json:"organization"`
	}

	userIdentifier := chi.URLParam(r, "user")
	userID, err := uuid.Parse(userIdentifier)
	if err != nil {
		userID, err = apiCfg.DB.GetUserIDfromName(r.Context(), userIdentifier)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting the user: %s", err))
			return
		}
	}

	userPublicInfo, err := apiCfg.DB.GetUserInfo(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting the user: %s", err))
		return
	}

	userPublicInfoJson := UserPublicInfo{
		ID:           userPublicInfo.ID,
		CreatedAt:    userPublicInfo.CreatedAt,
		Name:         userPublicInfo.Name,
		DisplayName:  userPublicInfo.DisplayName,
		Description:  userPublicInfo.Description,
		Verified:     userPublicInfo.Verified,
		Organization: userPublicInfo.Organization,
	}

	respondWithJSON(w, http.StatusOK, userPublicInfoJson)
}
