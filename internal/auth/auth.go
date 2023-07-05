package auth

import (
	"errors"
	"net/http"
	"strings"
)

/**GET API KEY FROM REQUEST HEADER */
func GetApiKey(headers http.Header) (string, error) {
	authKey := headers.Get("Authorization")

	if authKey == "" {
		return "", errors.New("couldn't find authorization key")
	}

	splittedKey := strings.Split(authKey, " ")
	if len(splittedKey) != 2 {
		return "", errors.New("malformed authorization key")
	}

	if splittedKey[0] != "ApiKey" {
		return "", errors.New("wrong authorization type")
	}

	return splittedKey[1], nil
}
