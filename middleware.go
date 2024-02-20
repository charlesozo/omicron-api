package main

import (
	"fmt"
	"github.com/charlesozo/omicron-api/internal/auth"
	"github.com/charlesozo/omicron-api/internal/database"
	"net/http"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.Registereduser)

func (cfg *ApiConfig) MiddleWareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetHeaderToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid API key- %s", err))
			return
		}
		if !auth.ValidateJWT(token, cfg.SecretKey) {
			respondWithError(w, http.StatusRequestTimeout, "Token has expired")
			return
		}
		user, err := cfg.DB.GetUserByToken(r.Context(), token)
		if err != nil {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("User doesn't exist in the database, %v", err))
			return
		}

		handler(w, r, user)
	}
}
