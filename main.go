package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book structs
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"ispn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
		json.NewEncoder(w).Encode(&Book{})
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000)) // convert int to str. Mock ID is not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(&Book{})
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000)) // convert int to str. Mock ID is not safe
			books = append(books, book)
			json.NewEncoder(w).Encode(&Book{})
			return
		}
	}
	json.NewEncoder(w).Encode(&books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(&books)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "1312", Title: "Book1", Author: &Author{Firstname: "John", Lastname: "James"}})
	books = append(books, Book{ID: "2", Isbn: "4145", Title: "Book2", Author: &Author{Firstname: "Sam", Lastname: "Smith"}})

	// Router handlers
	// every router handler function has to take response and request
	r.HandleFunc("/api/books", getBooks)
	r.HandleFunc("/api/books/{id}", getBook)
	r.HandleFunc("/api/books", createBook)
	r.HandleFunc("/api/books/{id}", updateBook)
	r.HandleFunc("/api/books/{id}", deleteBook)

	log.Fatal(http.ListenAndServe(":8000", r))
}
