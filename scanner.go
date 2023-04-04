package jsonquerier

import (
	"strings"
	"unicode"
)

const punctuations = "{}[]:,"

type Scanner struct {
	input string
	pos   int
}

func newScanner(input string) Scanner {
	return Scanner{input, 0}
}

func (s *Scanner) SkipSpaces() {
	for s.pos < len(s.input) && unicode.IsSpace(rune(s.input[s.pos])) {
		s.pos++
	}
}

func (s *Scanner) NextNumber() token {
	number := ""
	c := s.input[s.pos]
	for s.pos < len(s.input) && (unicode.IsDigit(rune(c)) || c == '.' || c == 'E' || c == '+' || c == '-') {
		number += string(c)
		s.pos++
		if s.pos < len(s.input) {
			c = s.input[s.pos]
		}
	}
	return newToken(number)
}

func (s *Scanner) NextString() token {
	str := ""

	c := s.input[s.pos]
	for s.pos < len(s.input) && (c != '"' || len(str) == 0) {
		str += string(c)
		s.pos++
		if s.pos < len(s.input) {
			c = s.input[s.pos]
		}
	}
	if c == '"' {
		str += "\""
		s.pos++
	}
	return newToken(str)
}

func (s *Scanner) NextToken() token {
	if s.pos >= len(s.input) {
		return token{"", EOF}
	}
	s.SkipSpaces()
	currentChar := rune(s.input[s.pos])
	if unicode.IsDigit(currentChar) || currentChar == '-' {
		return s.NextNumber()
	}
	if currentChar == '"' {
		return s.NextString()
	}
	currentSpelling := string(s.input[s.pos])

	s.pos++
	if strings.Contains(punctuations, currentSpelling) {
		return newToken(currentSpelling)
	}

	for s.pos < len(s.input) && !unicode.IsSpace(rune(s.input[s.pos])) {
		if strings.Contains(punctuations, string(s.input[s.pos])) {
			return newToken(currentSpelling)
		}
		currentSpelling += string(s.input[s.pos])
		s.pos++
	}

	return newToken(currentSpelling)
}
