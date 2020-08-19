package routers

import (
	"fmt"
	"net/http"

	"github.com/nahumsa/BooksAPI/books"

	"github.com/gin-gonic/gin"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "7894561230"
	dbname   = "books"
)

func dbInfo() string {
	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	PsqlInfo = fmt.Sprintf("%s dbname=%s", PsqlInfo, dbname)
	return PsqlInfo
}

// FindBooks prints all books upon a RestAPI request
func FindBooks(c *gin.Context) {
	PsqlInfo := dbInfo()
	db, err := books.Open("postgres", PsqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	book, err := db.AllBooks()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBookInput struct to validate input of data.
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// CreateBook creates a book upon a RestAPI request
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := books.Book{Title: input.Title, Author: input.Author}

	PsqlInfo := dbInfo()
	db, err := books.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Add(book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
