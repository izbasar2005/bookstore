package handlers

import (
	"api/models"
	"encoding/json"
	"net/http"
)

var authors []models.Author
var authorIDCounter = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authors)
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var newAuthor models.Author
	json.NewDecoder(r.Body).Decode(&newAuthor)

	if newAuthor.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newAuthor.ID = authorIDCounter
	authorIDCounter++
	authors = append(authors, newAuthor)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAuthor)
}
