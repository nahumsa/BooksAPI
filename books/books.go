package books

import (
	"database/sql"

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
		id INT NOT NULL,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL
	)`
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
