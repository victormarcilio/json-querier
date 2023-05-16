package jsonquerier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMultipleStringField(t *testing.T) {
	payload := `
		{
			"Name" : "Bob",
			"HasKids": false,
			"Married" : true,
			"Age" : 15,
			"Weigth" : 75.5e-105,
			"Address": null
		}`
	keys := []string{"Name", "HasKids", "Married", "Age", "Weigth", "Address"}
	expected := createMapWithExpectedKeys(keys)
	got := parse(payload)
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
	keys := []string{"Book", "Book.Title", "Book.Pages", "Book.Author", "Book.Author.Name", "Book.Author.Age"}
	expected := createMapWithExpectedKeys(keys)
	got := parse(payload)
	require.Equal(t, expected, got)
}

func TestParseArraysWithSimpleValues(t *testing.T) {
	payload := `
		{
			"Colors": ["blue", "green", "black"],
			"Numbers": [5, 3.5],
			"Empty": [],
			"Values": [true, false, null]
		}`
	keys := []string{"Colors", "Colors[0]", "Colors[1]", "Colors[2]", "Numbers", "Numbers[0]", "Numbers[1]",
		"Empty", "Values", "Values[0]", "Values[1]", "Values[2]"}

	expected := createMapWithExpectedKeys(keys)
	got := parse(payload)
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
	keys := []string{"Fruits", "Fruits[0]", "Fruits[0].Name", "Fruits[0].Price",
		"Fruits[1]", "Fruits[1].Name", "Fruits[1].Color"}
	expected := createMapWithExpectedKeys(keys)
	got := parse(payload)
	require.Equal(t, expected, got)
}

func TestParseArraysWithInnerArrays(t *testing.T) {
	payload := `
		{
			"Random": [
				["banana", "grapes", "apple"],
				[15, null, false, true, 25.5]
			]
		}`
	keys := []string{"Random", "Random[0]", "Random[0][0]", "Random[0][1]", "Random[0][2]",
		"Random[1]", "Random[1][0]", "Random[1][1]", "Random[1][2]", "Random[1][3]", "Random[1][4]"}
	expected := createMapWithExpectedKeys(keys)
	got := parse(payload)
	require.Equal(t, expected, got)
}

func TestParseCreateQuerier(t *testing.T) {
	payload := `
		{
			"Name" : "Bob",
			"Children": [
				{
					"Name": "Katty",
					"Age": 5
				},
				{
					"Name": "Fred",
					"Age": 3
				}
			],
			"Pets": ["Dog", "Cat"] 
		}`
	keys := []string{"Name", "Children", "Children[0]", "Children[0].Name", "Children[0].Age",
		"Children[1]", "Children[1].Name", "Children[1].Age", "Pets", "Pets[0]", "Pets[1]"}
	expected := createMapWithExpectedKeys(keys)
	got := parse(payload)
	require.Equal(t, expected, got)

	querier := CreateQuerier(payload)
	for _, key := range keys {
		require.True(t, querier(key))
	}
}

func createMapWithExpectedKeys(keys []string) map[string]bool {
	m := make(map[string]bool)
	for _, key := range keys {
		m[key] = true
	}
	return m
}
