package main

import (
	"log"
	"net/http"

	"api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.GetBookByID).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.DeleteBook).Methods("DELETE")

	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors", handlers.AddAuthor).Methods("POST")

	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	r.HandleFunc("/categories", handlers.AddCategory).Methods("POST")

	log.Println("Сервер запущен на порт :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
