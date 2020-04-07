package api

import (
	"encoding/json"
	"net/http"
)

// Book type
type Book struct {
	Title  string
	Author string
	ISBN   string
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
	return Book{}
}

//
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

}
