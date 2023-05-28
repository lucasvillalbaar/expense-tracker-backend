package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
	"github.com/lucasvillalbaar/expense-tracker-backend/repository"
	"github.com/segmentio/ksuid"
)

type CreateTransactionRequest struct {
	CreatedAt      time.Time `json:"created_at"`
	Type           string    `json:"type"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	Account        string    `json:"account"`
	OriginalAmount float64   `json:"original_amount"`
	Currency       string    `json:"currency"`
}

type UpdateTransactionRequest struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Type           string    `json:"type"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	Account        string    `json:"account"`
	OriginalAmount float64   `json:"original_amount"`
	Currency       string    `json:"currency"`
	BaseAmount     float64   `json:"base_amount"`
}

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdAt := time.Now().UTC()

	id, err := ksuid.NewRandom()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tx := transaction.Transaction{
		ID:             id.String(),
		CreatedAt:      createdAt,
		Type:           req.Type,
		Category:       req.Category,
		Description:    req.Description,
		Account:        req.Account,
		OriginalAmount: req.OriginalAmount,
		Currency:       req.Currency,
	}

	if err := repository.InsertTransaction(r.Context(), &tx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tx)
}

func DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionID := params["id"]

	if err := repository.DeleteTransaction(r.Context(), transactionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx := transaction.Transaction{
		ID:             req.ID,
		CreatedAt:      req.CreatedAt,
		Type:           req.Type,
		Category:       req.Category,
		Description:    req.Description,
		Account:        req.Account,
		OriginalAmount: req.OriginalAmount,
		Currency:       req.Currency,
		BaseAmount:     req.BaseAmount,
	}

	if err := repository.UpdateTransaction(r.Context(), &tx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tx)
}
