package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

var ErrNoAuthHeaderIncluded = errors.New("no auth header included in request")

type MyCustomsClaims struct {
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("invalid password")
	}
	dat, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(dat), err
}
func ComparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func MakeJWT(userID uuid.UUID, secrete_key string, expiresin time.Duration, issuer string) (string, error) {
	signingKey := []byte(secrete_key)
	claims := MyCustomsClaims{
		jwt.RegisteredClaims{
			Issuer:    string(issuer),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresin)),
			Subject:   fmt.Sprintf("%v", userID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}
func ValidateJWT(token_string, secret_key string) bool {
	token, err := jwt.ParseWithClaims(
		token_string,
		&MyCustomsClaims{},
		func(token *jwt.Token) (interface{}, error) { return []byte(secret_key), nil },
	)
	if err != nil {
		return false
	}
	return token.Valid
}

func GenerateAPIKey(email, salt string) (string, error) {
	if email == "" {
		return "", errors.New("email cannot be empty")
	}

	data := email + salt
	hash := sha256.New()
	hash.Write([]byte(data))
	apiKey := hex.EncodeToString(hash.Sum(nil))

	return apiKey, nil
}
func GetHeaderToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
