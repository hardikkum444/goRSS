package main

import (
	"fmt"
	"net/http"

	"github.com/hardikkum444/goRSS/internal/auth"
	"github.com/hardikkum444/goRSS/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		handler(w, r, user)
	}
}
