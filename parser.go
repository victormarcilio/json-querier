package jsonquerier

import (
	"fmt"
	"strings"
)

type parser struct {
	scanner
	fields       map[string]bool
	spelling     string
	currentToken token
}

func isSimpleValue(tk tokenID) bool {
	return tk == FALSE || tk == TRUE || tk == NUMBER || tk == STRING || tk == NULL
}

func (p *parser) parseObject() {
	p.SkipSpaces()
	p.acceptToken(OPEN_CURLY)
	p.acceptToken(STRING)
	p.appendStr()
	p.acceptToken(SEMICOLON)
	next := p.PeekNextToken()

	if isSimpleValue(next.ID) {
		p.acceptIt()
	} else if next.ID == OPEN_CURLY {
		p.parseObject()
	} else if next.ID == OPEN_SQUARE {
		p.parseArray()
	}

	for next = p.PeekNextToken(); next.ID == COMMA; next = p.PeekNextToken() {
		p.acceptIt()
		p.removeStr()
		p.acceptToken(STRING)
		p.appendStr()
		p.acceptToken(SEMICOLON)
		next = p.PeekNextToken()
		if isSimpleValue(next.ID) {
			p.acceptIt()
		} else if next.ID == OPEN_CURLY {
			p.parseObject()
		} else if next.ID == OPEN_SQUARE {
			p.parseArray()
		}
	}
	p.removeStr()
	p.acceptToken(CLOSE_CURLY)
}

func (p *parser) parseArray() {
	p.acceptToken(OPEN_SQUARE)
	next := p.PeekNextToken()
	if next.ID == CLOSE_SQUARE {
		p.acceptIt()
		return
	}
	spelling := p.spelling
	index := 0
	p.spelling = fmt.Sprintf("%s[%d]", spelling, index)
	p.addCurrentSpelling()
	if isSimpleValue(next.ID) {
		p.acceptIt()
	} else if next.ID == OPEN_CURLY {
		p.parseObject()
	} else if next.ID == OPEN_SQUARE {
		p.parseArray()
	}
	for next = p.PeekNextToken(); next.ID == COMMA; next = p.PeekNextToken() {
		p.acceptIt()
		next = p.PeekNextToken()
		index++
		p.spelling = fmt.Sprintf("%s[%d]", spelling, index)
		p.addCurrentSpelling()
		if isSimpleValue(next.ID) {
			p.acceptIt()
		} else if next.ID == OPEN_CURLY {
			p.parseObject()
		} else if next.ID == OPEN_SQUARE {
			p.parseArray()
		}
	}

	p.acceptToken(CLOSE_SQUARE)
}

func (p *parser) acceptToken(expected tokenID) {
	currentToken := p.NextToken()
	if currentToken.ID != expected {
		panic(fmt.Errorf("expected %v got %v", expected, currentToken.ID))
	}
	p.currentToken = currentToken
}

func (p *parser) acceptIt() {
	p.currentToken = p.NextToken()
}

func (p *parser) addCurrentSpelling() {
	p.fields[p.spelling] = true
}

func parse(input string) map[string]bool {
	p := parser{scanner: newScanner(input), fields: make(map[string]bool)}

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

func CreateQuerier(input string) func(string) bool {
	existing := parse(input)
	return func(field string) bool {
		return existing[field]
	}
}
