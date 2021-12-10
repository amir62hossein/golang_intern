package main

import (
	

	"encoding/json"
	
	"internship/db"
	"internship/models"
	_ "math/rand"
	"net/http"
	_ "strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var BooksDB []models.BooksDB

// get all books
func BookHandler(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	AllBooks := db.QueryAllBooks()

	json.NewEncoder(rw).Encode(AllBooks)

}

// get one book by id
func BookHandlerById(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	book := db.QuerySingelBook(id)

	json.NewEncoder(rw).Encode(book)

	book = nil

}

//delete one book
func DeleteBookHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	BooksDB = db.QueryDeleteBook(id)

	json.NewEncoder(rw).Encode(BooksDB)

}

// create one book
func CreateBookHandler(rw http.ResponseWriter, r *http.Request) {

	var bookName string

	json.NewDecoder(r.Body).Decode(&bookName)

	BooksDB = db.QueryCreateBook(bookName)

	json.NewEncoder(rw).Encode(BooksDB)

}

//edit book handler
func EditBookHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var newName string

	json.NewDecoder(r.Body).Decode(&newName)

	BooksDB = db.QueryEditBook(id , newName)

	json.NewEncoder(rw).Encode(BooksDB)

}
func main() {

	db, _ := db.ConnectDB()

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/books", BookHandler).Methods("GET")
	router.HandleFunc("/books/{id}", BookHandlerById).Methods("GET")
	router.HandleFunc("/books/{id}", DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/books", CreateBookHandler).Methods("POST")
	router.HandleFunc("/books/{id}", EditBookHandler).Methods("PUT")

	http.ListenAndServe(":8000", router)
}
