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

// dbInfo return a string that gives a route
// for the database.
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

	psqlInfo := dbInfo()
	db, err := books.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Add book and get its id
	id, err := db.Add(book)
	book.ID = id

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// FindOneBook finds a book given an id
func FindOneBook(c *gin.Context) {

	psqlInfo := dbInfo()
	db, err := books.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	book, err := db.FindBook(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
