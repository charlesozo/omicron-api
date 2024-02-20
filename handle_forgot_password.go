package main

import (
	"encoding/json"
	"fmt"
	"github.com/charlesozo/omicron-api/internal/auth"
	"github.com/charlesozo/omicron-api/internal/database"
	"net/http"
)

func (cfg *ApiConfig) handleForgotPassword(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	params := ForgotPasswordParams{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling json, %v", err))
		return
	}
	user, err := cfg.DB.GetUserDetails(r.Context(), database.GetUserDetailsParams{
		Email:          params.WhatsappNumber,
		WhatsappNumber: params.WhatsappNumber,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("user not found %v", err))
		return
	}
	apiKey, err := auth.GenerateAPIKey(user.Email, cfg.SecretKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	if err = sendMail("./templates/forgotpassword.html", user.Email, apiKey); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	respondWithJSON(w, http.StatusOK, "successfully sent forgot password mail")
}
