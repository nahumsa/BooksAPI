package main

import (
	"flag"
	"fmt"

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

	r.GET("/books", routers.FindBooks)
	r.POST("/books", routers.CreateBook)
	r.GET("/books/:id", routers.FindOneBook)
	r.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
