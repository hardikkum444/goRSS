package auth

import (
	"errors"
	"net/http"
	"strings"
)

func getAPIKey(headers http.Header) (string, error) {

	val := headers.Get("Authorisation")

	if val == "" {
		return "", errors.New("no authorisation info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of token")
	}

	return vals[1], nil

}