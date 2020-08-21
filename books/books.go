package books

import (
	"database/sql"
	"fmt"
	"strconv"

	// Import for running
	_ "github.com/lib/pq"
)

// Book struct with ID, Title and Author
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// DataBase is a struct that keeps a sql.DB
type DataBase struct {
	db *sql.DB
}

// Close Closes the database
func (db *DataBase) Close() error {
	return db.db.Close()
}

// AllBooks returns all books on the database
func (db *DataBase) AllBooks() ([]Book, error) {
	rows, err := db.db.Query("SELECT id, title, author FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		ret = append(ret, b)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

// FindBook returns a book from an id
func (db *DataBase) FindBook(id string) (Book, error) {
	b, err := findID(db.db, id)
	return b, err
}

func findID(db *sql.DB, id string) (Book, error) {
	statement := `SELECT title, author FROM book
				  WHERE id=$1`
	var b Book
	err := db.QueryRow(statement, id).Scan(&b.Title, &b.Author)
	b.ID, _ = strconv.Atoi(id)
	return b, err
}

// FindAuthor adds a input to the database
func (db *DataBase) FindAuthor(author string) ([]Book, error) {
	fmt.Println(author)
	statement := `SELECT id, title, author FROM book
				  WHERE author=$1`
	rows, err := db.db.Query(statement, author)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		ret = append(ret, b)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

// Add adds a input to the database
func (db *DataBase) Add(b Book) (int, error) {
	id, err := insertBook(db.db, b)
	return id, err
}

func insertBook(db *sql.DB, b Book) (int, error) {
	statement := `INSERT INTO book (title, author) 
				  VALUES ($1,$2) RETURNING id`
	var id int
	err := db.QueryRow(statement, b.Title, b.Author).Scan(&id)
	return id, err
}

// DeleteBook deletes a book entry given an id
func (db *DataBase) DeleteBook(b Book) error {
	return delete(db.db, b.ID)
}

func delete(db *sql.DB, id int) error {
	statement := `DELETE FROM book
				  WHERE id=$1`
	err := db.QueryRow(statement, id).Scan()
	return err
}

// Open Opens the database and returns a pointer to the database
func Open(driverName, dataSource string) (*DataBase, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DataBase{db}, nil
}

// Migrate Open the database and creates a book table
func Migrate(driverName, dataSource string) error {

	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = createBookTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func createBookTable(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS book (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL
	);`
	_, err := db.Exec(statement)
	return err
}

// Reset Resets the database
func Reset(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}

	return db.Close()
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}

	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)

	if err != nil {
		return err
	}

	return nil
}
