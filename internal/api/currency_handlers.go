package api

import (
	"encoding/json"
	"net/http"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
	"github.com/lucasvillalbaar/expense-tracker-backend/repository"
)

type CreateCurrencyRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func CreateCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateCurrencyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	curr := transaction.Currency{
		ID:     req.ID,
		Name:   req.Name,
		Symbol: req.Symbol,
	}

	if err := repository.InsertCurrency(r.Context(), &curr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&curr)
}
