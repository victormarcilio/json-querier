package token

import "strconv"

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
	INVALID
)

type token struct {
	spelling string
	ID       tokenID
}

func newToken(spelling string) token {
	switch spelling {
	case "{":
		return token{"{", OPEN_CURLY}
	case "}":
		return token{"}", CLOSE_CURLY}
	case "[":
		return token{"[", OPEN_SQUARE}
	case "]":
		return token{"]", CLOSE_SQUARE}
	case ",":
		return token{",", COMMA}
	case ":":
		return token{":", SEMICOLON}
	case "true":
		return token{"true", TRUE}
	case "false":
		return token{"false", FALSE}
	case "null":
		return token{"null", NULL}
	}

	if _, err := strconv.ParseFloat(spelling, 64); err == nil {
		return token{spelling, NUMBER}
	}
	return token{spelling, INVALID}
}
