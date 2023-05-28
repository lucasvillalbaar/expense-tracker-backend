package api

import (
	"encoding/json"
	"net/http"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
	"github.com/lucasvillalbaar/expense-tracker-backend/repository"
)

type CreateAccountRequest struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"` // Cash | Bank Account | Investment | Savings | Cripto Account
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account := transaction.Account{
		ID:       req.ID,
		Type:     req.Type,
		Name:     req.Name,
		Balance:  req.Balance,
		Currency: req.Currency,
	}

	if err := repository.InsertAccount(r.Context(), &account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&account)
}
