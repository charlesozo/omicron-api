package main

import (
	"database/sql"
	"github.com/charlesozo/omicron-api/internal/database"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type ApiConfig struct {
	Port      string
	DB_URL    string
	SecretKey string
	DB        *database.Queries
}

func main() {
	port, dburl, secretkey, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal("Unable to Connect to database", err)
	}
	queries := database.New(conn)
	cfg := ApiConfig{
		Port:      port,
		DB_URL:    dburl,
		SecretKey: secretkey,
		DB:        queries,
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", cfg.handleReadiness)
	v1Router.Post("/signup", cfg.handleUserSignUp)
	v1Router.Post("/confirmemail", cfg.handleConfirmEmail)
	v1Router.Post("/login", cfg.handleUserLogin)
	v1Router.Post("/update_username", cfg.MiddleWareAuth(cfg.handleUpdateUsername))
	v1Router.Post("/check_password", cfg.MiddleWareAuth(cfg.checkOldPassword))
	v1Router.Post("/update_password", cfg.MiddleWareAuth(cfg.handleUpdateOldPassword))
	v1Router.Post("/forgot_password", cfg.handleForgotPassword)

	router.Mount("/api/v1", v1Router)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
		Timeout: 
	}
	log.Printf("Server started at port %s", port)
	log.Fatal(server.ListenAndServe())
}
