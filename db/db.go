package db

import (
	"database/sql"

	"fmt"
	_ "github.com/lib/pq"
	"internship/models"
)

var BooksDB []models.BooksDB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nourian1999"
	dbname   = "book_store"
)

func ConnectDB() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	return db, err

}

func QueryAllBooks() []models.BooksDB {

	BooksDB = nil

	select_query := `SELECT * FROM books ORDER BY id;`

	db, err := ConnectDB()

	if err != nil {
		panic(err)
	}
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

	return BooksDB

}

func QuerySingelBook(id string) []models.BooksDB {

	BooksDB = nil

	db, err := ConnectDB()

	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT * FROM books WHERE id=$1;`

	var book models.BooksDB

	row := db.QueryRow(sqlStatement, id)

	err = row.Scan(&book.ID, &book.Name)

	if err == nil {
		BooksDB = nil
		BooksDB = append(BooksDB, book)
	} else {
		panic(err)
	}

	return BooksDB

}
func QueryDeleteBook(id string) []models.BooksDB {
	BooksDB = nil

	db, err := ConnectDB()

	if err != nil {
		panic(err)
	}
	deleteSqlStatement := `DELETE FROM books WHERE id = $1;`

	_, err = db.Exec(deleteSqlStatement, id)

	if err != nil {
		panic(err)
	}

	BooksDB = QueryAllBooks()

	return BooksDB

}
func QueryCreateBook(bookName string) []models.BooksDB {

	BooksDB = nil

	db, err := ConnectDB()

	if err != nil {
		panic(err)
	}

	createBookStatement := `INSERT INTO books (name) VALUES ($1);`

	_, err = db.Exec(createBookStatement , string(bookName))
	if err != nil {
		panic(err)
	}

	BooksDB = QueryAllBooks()

	return BooksDB

}
