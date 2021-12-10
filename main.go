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

var Test struct {
	Name string `json:"name"`
}

var BooksDB []models.BooksDB

// home routes
func HomeHnadler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Home address...")
}

// get all books
func BookHandler(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	_, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	AllBooks := db.QueryAllBooks()

	json.NewEncoder(rw).Encode(AllBooks)

	AllBooks = nil

}

// get one book by id
func BookHandlerById(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	db , err := db.ConnectDB()

	if err != nil {
		panic(err)
	}


	sqlStatement := `SELECT * FROM books WHERE id=$1;`

	var book models.BooksDB

	row := db.QueryRow(sqlStatement, id["id"])

	err = row.Scan(&book.ID, &book.Name)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		BooksDB = append(BooksDB, book)
		json.NewEncoder(rw).Encode(BooksDB)
		BooksDB = nil
		return
	default:
		panic(err)
	}

}

//delete one book
func DeleteBookHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	db , err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	deleteSqlStatement := `DELETE FROM books WHERE id = $1;`

	_, err = db.Exec(deleteSqlStatement, id)

	if err != nil {
		panic(err)
	}
	select_query := `SELECT * FROM books ORDER BY id;`

	rows, err := db.Query(select_query)

	for rows.Next() {
		var u models.BooksDB

		err := rows.Scan(&u.ID, &u.Name)

		if err != nil {
			panic(err)
		}

		BooksDB = append(BooksDB, u)

	}

	if err != nil {
		panic(err)
	}

	json.NewEncoder(rw).Encode(BooksDB)

}

// create one book
func CreateBookHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	var n string = "finall"

	//_ = json.NewDecoder(r.Body).Decode(&n)

	json.NewEncoder(rw).Encode(n)

	sqlInsertStatement := `INSERT INTO books (name) VALUES($1);`

	_, err = db.Exec(sqlInsertStatement, n)

	if err != nil {
		panic(err)
	}

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
	router.HandleFunc("/", HomeHnadler).Methods("GET")
	router.HandleFunc("/books", BookHandler).Methods("GET")
	router.HandleFunc("/books/{id}", BookHandlerById).Methods("GET")
	router.HandleFunc("/books/{id}", DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/books", CreateBookHandler).Methods("POST")
	router.HandleFunc("/books/{id}", EditBookHandler).Methods("PUT")

	http.ListenAndServe(":8000", router)
}
