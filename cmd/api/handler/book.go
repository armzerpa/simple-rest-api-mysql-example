package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/api/model"
	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/repository"
	"github.com/gin-gonic/gin"
)

func NewBookHandler(db *sql.DB) *HandlerBook {
	return &HandlerBook{
		bookRepository: repository.NewRepository(db),
	}
}

type HandlerBook struct {
	bookRepository repository.Repo
}

func (h *HandlerBook) GetBooks(c *gin.Context) {
	books := h.bookRepository.GetAll()
	c.IndentedJSON(http.StatusOK, books)
}

func (h *HandlerBook) GetBookById(c *gin.Context) {
	id := c.Param("id")
	book := h.bookRepository.GetById(id)
	c.IndentedJSON(http.StatusOK, &book)
}

func (h *HandlerBook) DeleteBookById(c *gin.Context) {
	id := c.Param("id")
	result := h.bookRepository.DeleteById(id)
	c.IndentedJSON(http.StatusOK, "Book deleted: "+fmt.Sprintf("%t", result))
}

func (h *HandlerBook) CreateBook(c *gin.Context) {
	var bookToInsert model.Book
	error := c.BindJSON(&bookToInsert)
	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid body")
		return
	}

	book := h.bookRepository.Create(bookToInsert)
	c.IndentedJSON(http.StatusCreated, book)
}
