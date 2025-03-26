package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/models"
	"github.com/gorilla/mux"
)

var books []models.Book
var bookIDCounter = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	categoryFilter := r.URL.Query().Get("category")
	authorFilter := r.URL.Query().Get("author_id")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}

	var filteredBooks []models.Book

	for _, book := range books {
		if categoryFilter != "" {
			categoryID, err := strconv.Atoi(categoryFilter)
			if err != nil || book.CategoryID != categoryID {
				continue
			}
		}

		if authorFilter != "" {
			authorID, err := strconv.Atoi(authorFilter)
			if err != nil || book.AuthorID != authorID {
				continue
			}
		}

		filteredBooks = append(filteredBooks, book)
	}

	start := (page - 1) * limit
	end := start + limit
	if start > len(filteredBooks) {
		start = len(filteredBooks)
	}
	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredBooks[start:end])
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	json.NewDecoder(r.Body).Decode(&newBook)

	if newBook.Title == "" || newBook.Price <= 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newBook.ID = bookIDCounter
	bookIDCounter++
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updatedBook models.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = book.ID
			books[i] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
