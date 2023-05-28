package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/lucasvillalbaar/expense-tracker-backend/database"
	"github.com/lucasvillalbaar/expense-tracker-backend/internal/api"
	"github.com/lucasvillalbaar/expense-tracker-backend/repository"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/transactions", api.CreateTransactionHandler).Methods(http.MethodPost)
	router.HandleFunc("/transactions/{id}", api.DeleteTransactionHandler).Methods(http.MethodDelete)
	router.HandleFunc("/transactions", api.UpdateTransactionHandler).Methods(http.MethodPut)
	router.HandleFunc("/currencies", api.CreateCurrencyHandler).Methods(http.MethodPost)
	router.HandleFunc("/currencies/{id}", api.DeleteCurrencyHandler).Methods(http.MethodDelete)
	router.HandleFunc("/currencies", api.UpdateCurrencyHandler).Methods(http.MethodPut)
	router.HandleFunc("/accounts", api.CreateAccountHandler).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id}", api.DeleteAccountHandler).Methods(http.MethodDelete)
	router.HandleFunc("/accounts", api.UpdateAccountHandler).Methods(http.MethodPut)
	router.HandleFunc("/categories", api.CreateCategoryHandler).Methods(http.MethodPost)
	return
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	repo, err := database.NewPostgresRepository(addr)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
