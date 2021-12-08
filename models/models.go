package models



type BooksDB struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PostBook struct {
	Name string `json:"name"`
}