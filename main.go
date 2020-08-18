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

	db, err := books.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	// RestAPI
	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	// })

	r.GET("/books", routers.FindBooks)

	r.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
