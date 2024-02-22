package main

import (
	"database/sql"
	"github.com/charlesozo/omicron-api/internal/database"
	"github.com/google/uuid"
	"time"
)

type PaymentStatusEnum string
type SubscriptionStatusEnum string
type SubscriptionTierEnum string

const (
	PaymentStatusEnumSuccessful   PaymentStatusEnum      = "Successful"
	PaymentStatusEnumPending      PaymentStatusEnum      = "Pending"
	PaymentStatusEnumFailed       PaymentStatusEnum      = "Failed"
	SubscriptionStatusEnumActive  SubscriptionStatusEnum = "Active"
	SubscriptionStatusEnumExpired SubscriptionStatusEnum = "Expired"
)

type NullPaymentStatusEnum struct {
	PaymentStatusEnum PaymentStatusEnum
	Valid             bool // Valid is true if PaymentStatusEnum is not NULL
}
type NullSubscriptionStatusEnum struct {
	SubscriptionStatusEnum SubscriptionStatusEnum
	Valid                  bool // Valid is true if SubscriptionStatusEnum is not NULL
}
type NullSubscriptionTierEnum struct {
	SubscriptionTierEnum SubscriptionTierEnum
	Valid                bool // Valid is true if SubscriptionTierEnum is not NULL
}
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
type SubscriptionModel struct {
	SubscriptionID   uuid.UUID                  `json:"subscription_id"`
	Userid           uuid.NullUUID              `json:"user_id"`
	ExpiryDate       time.Time                  `json:"expiry_date"`
	WhatsappNumber   string                     `json:"whatsapp_number"`
	Status           NullSubscriptionStatusEnum `json:"status"`
	RegisteredUser   sql.NullBool               `json:"registered_user"`
	SubscriptionTier NullSubscriptionTierEnum   `json:"subscription_tier"`
}

func dbSubModelToSubModel(sub database.Subscription) SubscriptionModel {
	return SubscriptionModel{
		SubscriptionID:   sub.SubscriptionID,
		Userid:           sub.Userid,
		ExpiryDate:       sub.ExpiryDate,
		WhatsappNumber:   sub.WhatsappNumber,
		Status:           NullSubscriptionStatusEnum{SubscriptionStatusEnum: SubscriptionStatusEnum(sub.Status.SubscriptionStatusEnum)},
		RegisteredUser:   sub.RegisteredUser,
		SubscriptionTier: NullSubscriptionTierEnum{SubscriptionTierEnum: SubscriptionTierEnum(sub.SubscriptionTier.SubscriptionTierEnum)},
	}
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
