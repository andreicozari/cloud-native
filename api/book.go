package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Book type
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// Marshall the book to json:
func (b Book) ToJSON() []byte {
	ToJson, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJson
}

func FromJSON(data []byte) Book {
	// define an empty book:
	book := Book{}
	err := json.Unmarshal(data, &book)

	if err != nil {
		panic(err)
	}

	return book
}

// Define a slice of books:
var books = []Book{
	Book{Title: "Cloud native GO", Author: "Rob Pike", ISBN: "0123456789"},
	Book{Title: "Book2", Author: "Rob Pike", ISBN: "1123456789"},
}

// Books handler to be used as http.HandleFunc for Book API:
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Info("Marshalling the books list: ", books)
	data, err := json.Marshal(books)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}
