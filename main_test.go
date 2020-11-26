package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestHomepageStatus(t *testing.T) {
	// Setup database and Grab our router
	setupTestDatabase(false)
	router := setupRouter()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBooksStatus(t *testing.T) {
	// Setup database and Grab our router
	setupDatabase(false)
	router := setupRouter()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/books")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestPOSTBook tests the post method for adding a book
func TestPOSTBook(t *testing.T) {

	// Expected response
	body := `{"data":{"id":1,"title":"Gödel, Escher, Bach: an Eternal Golden Braid","author":"Douglas Hofstader"}}`
	w := httptest.NewRecorder()

	// Setup database and Grab our router
	setupDatabase(true)
	router := setupRouter()

	values := map[string]string{"title": "Gödel, Escher, Bach: an Eternal Golden Braid", "author": "Douglas Hofstader"}

	jsonValue, _ := json.Marshal(values)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	readBuf, _ := ioutil.ReadAll(w.Body)
	assert.True(t, string(readBuf) == body)

}

func TestBookSearch(t *testing.T) {
	// Setup database and Grab our router
	setupDatabase(false)
	router := setupRouter()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/books/id/1")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Test body response
	body := `{"data":{"id":1,"title":"Gödel, Escher, Bach: an Eternal Golden Braid","author":"Douglas Hofstader"}}`
	readBuf, _ := ioutil.ReadAll(w.Body)
	assert.True(t, string(readBuf) == body)

}

func TestAuthorSearch(t *testing.T) {
	// Setup database and Grab our router
	setupDatabase(false)
	router := setupRouter()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/books/author/Douglas Hofstader")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Test body response
	body := `{"data":[{"id":1,"title":"Gödel, Escher, Bach: an Eternal Golden Braid","author":"Douglas Hofstader"}]}`
	readBuf, _ := ioutil.ReadAll(w.Body)
	assert.True(t, string(readBuf) == body)

}

// func TestDELETEBook(t *testing.T) {

// 	// Expected response
// 	body := `{"data":{"id":1,"title":"Gödel, Escher, Bach: an Eternal Golden Braid","author":"Douglas Hofstader"}}`
// 	w := httptest.NewRecorder()

// 	// Setup database and Grab our router
// 	setupDatabase(true)
// 	router := setupRouter()

// 	values := map[string]string{"title": "Gödel, Escher, Bach: an Eternal Golden Braid", "author": "Douglas Hofstader"}

// 	jsonValue, _ := json.Marshal(values)
// 	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("X-Custom-Header", "myvalue")
// 	req.Header.Set("Content-Type", "application/json")

// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	readBuf, _ := ioutil.ReadAll(w.Body)
// 	assert.True(t, string(readBuf) == body)

// }
