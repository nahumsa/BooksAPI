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
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "7894561230"
	dbname     = "books"
	dbnameTest = "test"
)

func main() {
	reset := flag.Bool("reset", false, "true if you want to reset your database")

	flag.Parse()

	setupDatabase(*reset)

	r := setupRouter()

	r.Run()

}

// setupDatabase resets the database if reset is true and migrates it
func setupDatabase(reset bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	if reset {
		fmt.Println("DB Reseted.")
		must(books.Reset("postgres", psqlInfo, dbname))
	}

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(books.Migrate("postgres", psqlInfo))
}

// setupTestDatabase resets the database and migrates for tests
func setupTestDatabase(reset bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	if reset {
		must(books.Reset("postgres", psqlInfo, dbnameTest))
	}

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbnameTest)
	must(books.Migrate("postgres", psqlInfo))
}

// setupRouter creates the routing of the Books API
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Books API",
			},
		)

	})

	router.GET("/books", routers.FindBooks)
	router.POST("/books", routers.CreateBook)
	router.GET("/books/id/:id", routers.FindOneBook)
	router.DELETE("/books/id/:id", routers.DeleteBook)
	router.GET("/books/author/:author", routers.FindAuthor)
	return router

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
