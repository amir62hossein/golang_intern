package main

import (
	"database/sql"

	"encoding/json"
	"fmt"
	"internship/db"
	"internship/models"
	_ "math/rand"
	"net/http"
	_ "strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nourian1999"
	dbname   = "book_store"
)

var BooksDB []models.BooksDB

// get all books
func BookHandler(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	_, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	AllBooks := db.QueryAllBooks()

	json.NewEncoder(rw).Encode(AllBooks)

}

// get one book by id
func BookHandlerById(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	_, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

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

	//id := mux.Vars(r)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	_, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	var n models.PostBook

	_ = json.NewDecoder(r.Body).Decode(&n.Name)

	json.NewEncoder(rw).Encode(n)

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
