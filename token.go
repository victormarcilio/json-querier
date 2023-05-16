package jsonquerier

import (
	"regexp"
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

var spellingToID = map[string]tokenID{
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

var strRegex = regexp.MustCompile("^\".*\"$")
var numberRegex = regexp.MustCompile(`^-?(\d|[1-9]\d*)(\.\d+)?([Ee][-+]?\d+)?$`)

func getID(spelling string) tokenID {
	if id, ok := spellingToID[spelling]; ok {
		return id
	}

	if numberRegex.Match([]byte(spelling)) {
		return NUMBER
	}

	if strRegex.Match([]byte(spelling)) {
		return STRING
	}
	return INVALID
}

func newToken(spelling string) token {

	return token{spelling, getID(spelling)}
}
