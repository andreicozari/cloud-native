package api

import (
	"encoding/json"
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

//
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

}
