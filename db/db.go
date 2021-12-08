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
