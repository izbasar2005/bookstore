package handlers

import (
	"api/models"
	"encoding/json"
	"net/http"
)

var categories []models.Category
var categoryIDCounter = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(categories)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	json.NewDecoder(r.Body).Decode(&newCategory)

	if newCategory.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newCategory.ID = categoryIDCounter
	categoryIDCounter++
	categories = append(categories, newCategory)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}
