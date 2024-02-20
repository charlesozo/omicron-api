package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/charlesozo/omicron-api/internal/auth"
	"github.com/charlesozo/omicron-api/internal/database"
)

func (cfg *ApiConfig) handleUpdateUsername(w http.ResponseWriter, r *http.Request, user database.Registereduser) {
	decode := json.NewDecoder(r.Body)
	params := UpdateUsername{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling json, %v", err))
		return
	}
	updateuser, err := cfg.DB.Updateusername(r.Context(), database.UpdateusernameParams{
		Username:        params.Username,
		UpdatedAt:       time.Now().UTC(),
		UpdatedUsername: sql.NullBool{Bool: true, Valid: true},
		ID:              user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error updating username")
		return
	}

	respondWithJSON(w, http.StatusOK, dbRegUserToRegUser(updateuser))
}

func (cfg *ApiConfig) checkOldPassword(w http.ResponseWriter, r *http.Request, user database.Registereduser) {
	decode := json.NewDecoder(r.Body)
	params := UpdatePassword{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling json, %v", err))
		return
	}
	if err = auth.ComparePassword(params.Password, user.Password); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	respondWithJSON(w, http.StatusOK, "user authorized")
}

func (cfg *ApiConfig) handleUpdateOldPassword(w http.ResponseWriter, r *http.Request, user database.Registereduser) {
	decode := json.NewDecoder(r.Body)
	params := UpdatePassword{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling json, %v", err))
		return
	}
	updatedUser, err := cfg.DB.Updatepassword(r.Context(), database.UpdatepasswordParams{
		Password:        params.Password,
		UpdatedAt:       time.Now().UTC(),
		UpdatedPassword: sql.NullBool{Bool: true, Valid: true},
		ID:              user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error updating password")
		return
	}
	respondWithJSON(w, http.StatusOK, dbRegUserToRegUser(updatedUser))
}
