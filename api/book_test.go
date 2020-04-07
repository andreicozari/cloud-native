package api

import (
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

	assert.Equal(t, `{"Title":"Cloud native GO","Author":"Rob Pike","ISBN":"0123456789"}`,
		string(json), "Wrong marshalling the book json")
}

func TestBookFromJSON(t *testing.T) {
	data := []byte(`{"Title":"Cloud native GO","Author":"Rob Pike","ISBN":"0123456789"}`)
	book := FromJSON(data)

	assert.Equal(t, book, Book{Title: "Cloud native GO", Author: "Rob Pike", ISBN: "0123456789"})
}
