package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahumsa/BooksAPI/books"
	"github.com/nahumsa/BooksAPI/routers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "7894561230"
	dbname   = "books"
)

func main() {
	reset := flag.Bool("reset", false, "true if you want to reset your database")
	flag.Parse()

	// Database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	if *reset {
		fmt.Println("DB Reseted.")
		must(books.Reset("postgres", psqlInfo, dbname))
	}

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(books.Migrate("postgres", psqlInfo))

	// RestAPI
	r := gin.Default()

	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Books API",
			},
		)

	})

	r.GET("/books", routers.FindBooks)
	r.POST("/books", routers.CreateBook)
	r.GET("/books/id/:id", routers.FindOneBook)
	r.DELETE("/books/id/:id", routers.DeleteBook)
	r.GET("/books/author/:author", routers.FindAuthor)

	r.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
