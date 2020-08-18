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

// FindBooks prints all books upon a RestAPI request
func FindBooks(c *gin.Context) {
	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	PsqlInfo = fmt.Sprintf("%s dbname=%s", PsqlInfo, dbname)
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
