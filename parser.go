package jsonquerier

import (
	"fmt"
)

type Scanner interface {
	SkipSpaces()
	NextNumber() token
	NextString() token
	NextToken() token
}

type parser struct {
	scanner      Scanner
	fields       map[string]bool
	spelling     string
	currentToken token
}

func (p *parser) parseObject() {
	p.scanner.SkipSpaces()
	p.acceptToken(OPEN_CURLY)
	p.acceptToken(STRING)
	p.appendStr()
	p.acceptToken(SEMICOLON)
	p.acceptToken(STRING)
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

func Parse(input string) map[string]bool {
	scn := newScanner(input)
	p := parser{scanner: &scn, fields: make(map[string]bool)}

	p.parseObject()
	return p.fields
}

func removeQuotes(input string) string {
	return input[1 : len(input)-1]
}

func (p *parser) appendStr() {
	spelling := removeQuotes(p.currentToken.spelling)
	if p.spelling == "" {
		p.spelling += spelling
	}
	p.fields[p.spelling] = true
}

func (p *parser) removeStr() {
	p.spelling = ""
}
