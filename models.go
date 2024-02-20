package main

import (
	"database/sql"
	"github.com/charlesozo/omicron-api/internal/database"
	"github.com/google/uuid"
	"time"
)

type SignUpParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	WhatsappNumber string `json:"whatsapp_number"`
}
type LoginInParams struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	WhatsappNumber string `json:"whatsapp_number"`
}
type ForgotPasswordParams struct {
	Email          string `json:"email"`
	WhatsappNumber string `json:"whatsapp_number"`
}
type UpdateUsername struct {
	Username string `json:"username"`
}
type UpdatePassword struct {
	Password string `json:"password"`
}
type NewUserModel struct {
	ID              uuid.UUID    `json:"id"`
	Username        string       `json:"username"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	Email           string       `json:"email"`
	Isemailverified sql.NullBool `json:"is_email_verified"`
	VerifiedAt      time.Time    `json:"verified_at"`
	Password        string       `json:"password"`
	WhatsappNumber  string       `json:"whatsapp_number"`
	Loggedin        sql.NullBool `json:"logged_in"`
	Apikey          string       `json:"api_key"`
	UpdatedUsername sql.NullBool `json:"updated_username"`
	UpdatedPassword sql.NullBool `json:"updated_password"`
}

func dbRegUserToRegUser(user database.Registereduser) NewUserModel {
	return NewUserModel{
		ID:              user.ID,
		Username:        user.Username,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		Email:           user.Email,
		Isemailverified: user.Isemailverified,
		VerifiedAt:      user.VerifiedAt,
		Password:        user.Password,
		WhatsappNumber:  user.WhatsappNumber,
		Loggedin:        user.Loggedin,
		Apikey:          user.Apikey,
		UpdatedUsername: user.UpdatedUsername,
		UpdatedPassword: user.UpdatedPassword,
	}
}
