package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
	"github.com/lucasvillalbaar/expense-tracker-backend/repository"
	"github.com/segmentio/ksuid"
)

type CreateCategoryRequest struct {
	ID              string `json:"id"`
	CategoryName    string `json:"category_name"`
	SubcategoryName string `json:"subcategory_name"`
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := ksuid.NewRandom()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category := transaction.Category{
		ID:              id.String(),
		CategoryName:    req.CategoryName,
		SubcategoryName: req.SubcategoryName,
	}

	if err := repository.InsertCategory(r.Context(), &category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&category)
}

type UpdateCategoryRequest struct {
	ID              string `json:"id"`
	CategoryName    string `json:"category_name"`
	SubcategoryName string `json:"subcategory_name"`
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID := mux.Vars(r)["id"]

	err := repository.DeleteCategory(r.Context(), categoryID)
	if err != nil {
		if err.Error() == "category not found" {
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category := transaction.Category{
		ID:              req.ID,
		CategoryName:    req.CategoryName,
		SubcategoryName: req.SubcategoryName,
	}

	err := repository.UpdateCategory(r.Context(), &category)
	if err != nil {
		if err.Error() == "category not found" {
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&category)
}
