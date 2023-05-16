package jsonquerier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSingleStringField(t *testing.T) {
	payload := `
		{
			"Name" : "Bob"
		}`
	expected := map[string]bool{"Name": true}
	got := Parse(payload)
	require.Equal(t, expected, got)
}

func TestParseMultipleStringField(t *testing.T) {
	payload := `
		{
			"Name" : "Bob",
			"HasKids": false,
			"Married" : true,
			"Age" : 15,
			"Weigth" : 75.5,
			"Address": null
		}`
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

func TestParseArraysWithSimpleValues(t *testing.T) {
	payload := `
		{
			"Colors": ["blue", "green", "black"],
			"Numbers": [5, 3.5],
			"Values": [true, false, null]
		}`
	expected := map[string]bool{"Colors": true, "Colors[0]": true, "Colors[1]": true, "Colors[2]": true, "Numbers": true, "Numbers[0]": true, "Numbers[1]": true, "Values": true, "Values[0]": true, "Values[1]": true, "Values[2]": true}
	got := Parse(payload)
	require.Equal(t, expected, got)
}

func TestParseArraysWithInnerObjects(t *testing.T) {
	payload := `
		{
			"Fruits": [
				{
					"Name": "banana",
					"Price": 3.5
				},
				{
					"Name": "grapes",
					"Color": "green"
				}
			]
		}`
	expected := map[string]bool{"Fruits": true, "Fruits[0]": true, "Fruits[0].Name": true, "Fruits[0].Price": true, "Fruits[1]": true, "Fruits[1].Name": true, "Fruits[1].Color": true}
	got := Parse(payload)
	require.Equal(t, expected, got)
}
