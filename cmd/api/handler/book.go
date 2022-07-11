package handler

import (
	"database/sql"
	"net/http"

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
	//TODO: implement method
	//id := c.Param("id")
	//result := h.bookRepository.DeleteById(id)
	c.IndentedJSON(http.StatusOK, "Book deleted")
}

func (h *HandlerBook) CreateBook(c *gin.Context) {
	//TODO: implement method
	c.IndentedJSON(http.StatusCreated, "Book created")
}
