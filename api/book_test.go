package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookToJSON(t *testing.T) {
	book := Book{
		Title:  "Cloud native GO",
		Author: "Rob Pike",
		ISBN:   "0123456789",
	}

	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud native GO","author":"Rob Pike","isbn":"0123456789"}`,
		string(json), "Wrong marshalling the book json")
}

func TestBookFromJSON(t *testing.T) {
	byteData := []byte(`{"title":"Cloud native GO","author":"Rob Pike","isbn":"0123456789"}`)
	book := FromJSON(byteData)

	newBook, err := NewBook("Cloud native GO", "Rob Pike", "0123456789")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	assert.Equal(t, book, *newBook)
}
