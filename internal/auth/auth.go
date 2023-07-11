package auth

import (
	"errors"
	"net/http"
	"strings"
)

/**GET API KEY FROM REQUEST HEADER */
func GetApiKey(r *http.Request) (string, error) {
	cookie, err := r.Cookie("Authorization")
	if err != nil {
		return "", errors.New("couldn't find authorization key")
	}
	authKey := cookie.Value

	if authKey == "" {
		return "", errors.New("couldn't find authorization key")
	}

	splittedKey := strings.Split(authKey, " ")
	if len(splittedKey) != 2 {
		return "", errors.New("malformed authorization key")
	}

	if splittedKey[0] != "ApiKey" {
		return "", errors.New("wrong authorization type, should be ApiKey <key>")
	}

	return splittedKey[1], nil
}
