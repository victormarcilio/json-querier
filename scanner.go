package jsonquerier

type Scanner struct {
	input string
	pos   int
}

func newScanner(input string) Scanner {
	return Scanner{input, 0}
}

func (s *Scanner) readNextToken() token {
	if s.pos >= len(s.input) {
		return token{"", EOF}
	}
	currentToken := string(s.input[s.pos])
	s.pos++
	return newToken(currentToken)
}
