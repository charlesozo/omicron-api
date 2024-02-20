package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv() (string, string, string, error) {
	log.Fatal(godotenv.Load(".env"))
	port := os.Getenv("PORT")
	if port == "" {
		return "", "", "", errors.New("port is not set")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return port, "", "", errors.New("dburl is not set")
	}
	secretKey := os.Getenv("SECRET_KEY")
	return port, dbURL, secretKey, nil

}
