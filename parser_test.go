package jsonquerier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSingleStringField(t *testing.T) {
	payload := `{"Name" : "Bob"}`
	expected := map[string]bool{"Name": true}
	got := Parse(payload)
	require.Equal(t, expected, got)
}

func TestParseMultipleStringField(t *testing.T) {
	payload := `{"Name" : "Bob", "HasKids": false, "Married" : true, "Age" : 15, "Weigth" : 75.5, "Address": null}`
	expected := map[string]bool{"Name": true, "HasKids": true, "Married": true, "Age": true, "Weigth": true, "Address": true}
	got := Parse(payload)
	require.Equal(t, expected, got)
}

func TestParseWithRecursiveObjects(t *testing.T) {
	payload := `
		{
			"Book":{
				"Title": "Untitled",
				"Pages": 300,
				"Author": {
					"Name": "John Smith",
					"Age": 35
				}
			}
		}`
	expected := map[string]bool{"Book": true, "Book.Title": true, "Book.Pages": true, "Book.Author": true, "Book.Author.Name": true, "Book.Author.Age": true}
	got := Parse(payload)
	require.Equal(t, expected, got)
}
