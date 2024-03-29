package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/charlesozo/omicron-api/internal/auth"
	"github.com/charlesozo/omicron-api/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *ApiConfig) handleUserSignUp(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	params := SignUpParams{}
	err := decode.Decode(&params)
	userid := uuid.New()
	subscriptionid := uuid.New()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling json, %v", err))
		return
	}
	if err := isValidEmail(params.Email); err != nil {
		respondWithError(w, http.StatusNotAcceptable, err.Error())
		return
	}

	if err := isStrongPassword(params.Password); err != nil {
		respondWithError(w, http.StatusNotAcceptable, err.Error())
		return
	}
	if !containsNumbersOrSymbols(params.Username) {
		respondWithError(w, http.StatusNotAcceptable, "username not acceptable,letters only")
		return
	}
	if !isWhatsappNumbers(params.WhatsappNumber, cfg.DB_URL) {
		respondWithError(w, http.StatusNotAcceptable, "number is not on whatsapp")
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error hashing password, %v", err))
		return
	}
	apiKey, err := auth.GenerateAPIKey(params.Email, cfg.SecretKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}

	newUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             userid,
		Username:       params.Username,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Email:          params.Email,
		Password:       hashedPassword,
		WhatsappNumber: params.WhatsappNumber,
		Apikey:         apiKey,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	//generate a subscription table for the new registered table.
	subs, err := cfg.DB.CreateUserSubscription(r.Context(), database.CreateUserSubscriptionParams{
		SubscriptionID: subscriptionid,
		Userid:         uuid.NullUUID{UUID: userid, Valid: true},
		WhatsappNumber: newUser.WhatsappNumber,
		ExpiryDate:     time.Now().AddDate(0, 0, 10),
		RegisteredUser: sql.NullBool{Bool: true, Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating sub table %v", err))
		return
	}
	//send confirmation email
	err = sendMail("./templates/confirmationmail.html", params.Email, apiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error sending email %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, dbSubModelToSubModel(subs))
	respondWithJSON(w, http.StatusOK, dbRegUserToRegUser(newUser))
}
