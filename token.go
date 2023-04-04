package jsonquerier

import (
	"regexp"
	"strconv"
)

type tokenID int

const (
	OPEN_CURLY tokenID = iota
	CLOSE_CURLY
	OPEN_SQUARE
	CLOSE_SQUARE
	COMMA
	SEMICOLON
	TRUE
	FALSE
	NULL
	STRING
	NUMBER
	EOF
	INVALID
)

type token struct {
	spelling string
	ID       tokenID
}

var getID = func() func(string) tokenID {

	spellingToID := map[string]tokenID{
		"{":     OPEN_CURLY,
		"}":     CLOSE_CURLY,
		"[":     OPEN_SQUARE,
		"]":     CLOSE_SQUARE,
		",":     COMMA,
		":":     SEMICOLON,
		"true":  TRUE,
		"false": FALSE,
		"null":  NULL,
	}

	strRegex := regexp.MustCompile("^\".*\"$")

	return func(spelling string) tokenID {
		if id, ok := spellingToID[spelling]; ok {
			return id
		}

		if _, err := strconv.ParseFloat(spelling, 64); err == nil {
			return NUMBER
		}

		if strRegex.Match([]byte(spelling)) {
			return STRING
		}
		return INVALID
	}
}()

func newToken(spelling string) token {

	return token{spelling, getID(spelling)}
}
