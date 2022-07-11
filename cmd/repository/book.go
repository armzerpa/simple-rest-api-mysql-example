package repository

import (
	"database/sql"
	"log"

	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/api/model"
)

var books = []model.Book{
	{ID: "1", Title: "The Lord of the Rings", Author: "Tolkien", Quantity: 1},
	{ID: "2", Title: "Harry Potter", Author: "Rowling", Quantity: 5},
	{ID: "3", Title: "Te Bible", Author: "Apostols", Quantity: 4},
}

type Repo interface {
	GetAll() []model.Book
	GetById(id string) *model.Book
}

type BookRepo struct {
	DbConnection *sql.DB
}

func NewRepository(db *sql.DB) Repo {
	return BookRepo{DbConnection: db}
}

func (b BookRepo) GetAll() []model.Book {
	rows, err := b.DbConnection.Query("SELECT * FROM book")
	if err != nil {
		log.Println("Error in the query to the database")
		return nil
	}
	defer rows.Close()

	var bks []model.Book

	for rows.Next() {
		var bk model.Book

		err := rows.Scan(&bk.ID, &bk.Title, &bk.Author, &bk.Quantity)
		if err != nil {
			log.Println("Some error scanning data from the database")
			return nil
		}

		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Println("Some error selecting data from the database")
		return nil
	}

	return bks
}

func (b BookRepo) GetById(id string) *model.Book {
	var book model.Book
	err := b.DbConnection.QueryRow("SELECT * FROM book WHERE ID = ?", id).Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)

	if err != nil {
		log.Println("Error in the query to the database ", err)
		return nil
	}
	return &book
}
