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

func (cfg *ApiConfig) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	params := LoginInParams{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling json, %v", err))
		return
	}
	user, err := cfg.DB.GetUserDetails(r.Context(), database.GetUserDetailsParams{
		Email:          params.Email,
		WhatsappNumber: params.WhatsappNumber,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "user not found")
		return
	}
	err = auth.ComparePassword(params.Password, user.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	duration := time.Hour * 2
	token, err := auth.MakeJWT(user.ID, cfg.SecretKey, duration, user.Username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}

	err = cfg.DB.UpdateApiKey(r.Context(), database.UpdateApiKeyParams{
		Apikey: token,
		ID:     user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}

	err = cfg.DB.UserSigninout(r.Context(), database.UserSigninoutParams{
		ID:       user.ID,
		Loggedin: sql.NullBool{Bool: true, Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to sign user in")
		return
	}

	respondWithJSON(w, http.StatusOK, dbRegUserToRegUser(user))
}
