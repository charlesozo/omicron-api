package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/charlesozo/omicron-api/internal/auth"
	"github.com/charlesozo/omicron-api/internal/database"
)

func (cfg *ApiConfig) handleConfirmEmail(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetHeaderToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid API key- %s", err))
		return
	}
	user, err := cfg.DB.GetUserByToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "user not found")
		return
	}
	err = cfg.DB.VerifyUserEmail(r.Context(), database.VerifyUserEmailParams{
		Isemailverified: sql.NullBool{Bool: true, Valid: true},
		VerifiedAt:      time.Now().UTC(),
		ID:              user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to verify email")
	}
	respondWithJSON(w, http.StatusOK, "email verified")
}
