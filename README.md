# Books API using golang

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/a4e37324ec5fd722f252)
[![Documentation](https://godoc.org/github.com/nahumsa/BooksAPI?status.svg)](https://godoc.org/github.com/nahumsa/BooksAPI)

## Introduction
This is an REST API that is constructed to read a database of books. This is based on the article from [Rahman Fadhil](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/), where I adapted the implementation to use PostgreSQL. I also made a front end in order to search for books from a given author or a given id.

## Usage
For the first time you run, use the `-reset` flag in order to create the database in your local computer. There is also a frontend that is on the `localhost:8080`.

## Testing
All tests were made on postman and are shared on this [link](https://app.getpostman.com/run-collection/a4e37324ec5fd722f252) or in the button above.



## Todo:

Things that I plan to implement on the future.

- [x] Create database
- [x] Migrate database
- [x] Add a flag to reset the database
- [x] Add books to the database
- [x] Delete books
- [x] Query all books
- [x] Query book id
- [x] Query selected author
- [x] Create a frontend
    - [x] Show all books
    - [x] Find a book using the ID    
    - [x] Find a book by the Author
- [x] Create tests
- [ ] Deploy on Heroku
