package jsonquerier

import (
	"fmt"
	"strings"
)

type Scanner interface {
	SkipSpaces()
	NextNumber() token
	NextString() token
	NextToken() token
	PeekNextToken() token
}

type parser struct {
	scanner      Scanner
	fields       map[string]bool
	spelling     string
	currentToken token
}

func isSimpleValue(tk tokenID) bool {
	return tk == FALSE || tk == TRUE || tk == NUMBER || tk == STRING || tk == NULL
}

func (p *parser) parseObject() {
	p.scanner.SkipSpaces()
	p.acceptToken(OPEN_CURLY)
	p.acceptToken(STRING)
	p.appendStr()
	p.acceptToken(SEMICOLON)
	next := p.scanner.PeekNextToken()

	if isSimpleValue(next.ID) {
		p.acceptIt()
	} else if next.ID == OPEN_CURLY {
		p.parseObject()
	}

	for next = p.scanner.PeekNextToken(); next.ID == COMMA; next = p.scanner.PeekNextToken() {
		p.acceptIt()
		p.removeStr()
		p.acceptToken(STRING)
		p.appendStr()
		p.acceptToken(SEMICOLON)
		next = p.scanner.PeekNextToken()
		if isSimpleValue(next.ID) {
			p.acceptIt()
		} else if next.ID == OPEN_CURLY {
			p.parseObject()
		}
	}
	p.removeStr()
	p.acceptToken(CLOSE_CURLY)
}

func (p *parser) acceptToken(expected tokenID) {
	currentToken := p.scanner.NextToken()
	if currentToken.ID != expected {
		panic(fmt.Errorf("expected %v got %v", expected, currentToken.ID))
	}
	p.currentToken = currentToken
}

func (p *parser) acceptIt() {
	p.currentToken = p.scanner.NextToken()
}

func Parse(input string) map[string]bool {
	scn := newScanner(input)
	p := parser{scanner: &scn, fields: make(map[string]bool)}

	p.parseObject()
	p.acceptToken(EOF)
	return p.fields
}

func removeQuotes(input string) string {
	return input[1 : len(input)-1]
}

func (p *parser) appendStr() {
	spelling := removeQuotes(p.currentToken.spelling)
	if p.spelling != "" {
		p.spelling += "."
	}
	p.spelling += spelling
	p.fields[p.spelling] = true
}

func (p *parser) removeStr() {
	if ind := strings.LastIndex(p.spelling, "."); ind != -1 {
		p.spelling = p.spelling[:ind]
	} else {
		p.spelling = ""
	}
}
