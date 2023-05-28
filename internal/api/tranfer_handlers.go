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

type CreateTransferRequest struct {
	CreatedAt          time.Time `json:"created_at"`
	SourceAccount      string    `json:"source_account"`
	SourceAmount       float64   `json:"source_amount"`
	SourceFee          float64   `json:"source_fee"`
	DestinationAccount string    `json:"destination_account"`
	DestinationAmount  float64   `json:"destination_amount"`
	DestinationFee     float64   `json:"destination_fee"`
}

type UpdateTransferRequest struct {
	ID                 string    `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	SourceAccount      string    `json:"source_account"`
	SourceAmount       float64   `json:"source_amount"`
	SourceFee          float64   `json:"source_fee"`
	DestinationAccount string    `json:"destination_account"`
	DestinationAmount  float64   `json:"destination_amount"`
	DestinationFee     float64   `json:"destination_fee"`
}

func CreateTransferHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTransferRequest

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

	transfer := transaction.Transfer{
		ID:                 id.String(),
		CreatedAt:          createdAt,
		SourceAccount:      req.SourceAccount,
		SourceAmount:       req.SourceAmount,
		SourceFee:          req.SourceFee,
		DestinationAccount: req.DestinationAccount,
		DestinationAmount:  req.DestinationAmount,
		DestinationFee:     req.DestinationFee,
	}

	if err := repository.InsertTransfer(r.Context(), &transfer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&transfer)
}

func DeleteTransferHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transferID := params["id"]

	if err := repository.DeleteTransfer(r.Context(), transferID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateTransferHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateTransferRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transfer := transaction.Transfer{
		ID:                 req.ID,
		CreatedAt:          req.CreatedAt,
		SourceAccount:      req.SourceAccount,
		SourceAmount:       req.SourceAmount,
		SourceFee:          req.SourceFee,
		DestinationAccount: req.DestinationAccount,
		DestinationAmount:  req.DestinationAmount,
		DestinationFee:     req.DestinationFee,
	}

	if err := repository.UpdateTransfer(r.Context(), &transfer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&transfer)
}
