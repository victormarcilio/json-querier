package jsonquerier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type scannerTestCase struct {
	input          string
	expectedTokens []tokenID
}

func TestScanner_NextTokenShouldCategorizeSingleCharacterTokens(t *testing.T) {
	singleCharTokens := []string{"{", "}", "[", "]", ",", ":"}

	for _, spelling := range singleCharTokens {
		scanner := newScanner(spelling)
		token := scanner.NextToken()
		expectedToken := newToken(spelling)
		require.Equal(t, expectedToken, token)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}

func TestScanner_NextTokenShouldCategorizeMultipleCharacterFixedTokens(t *testing.T) {
	multipleCharTokens := []string{"true", "false", "null"}

	for _, spelling := range multipleCharTokens {
		scanner := newScanner(spelling)
		token := scanner.NextToken()
		expectedToken := newToken(spelling)
		require.Equal(t, expectedToken, token)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}

func TestScanner_NextTokenShouldDetectInvalidTokens(t *testing.T) {
	multipleCharTokens := []string{"tru", "trues", "fals", "falses", "nul", "nully"}

	for _, spelling := range multipleCharTokens {
		scanner := newScanner(spelling)
		token := scanner.NextToken()
		require.Equal(t, INVALID, token.ID)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}

func TestScanner_ShouldSeparateSingleCharacterTokens(t *testing.T) {

	testcases := []scannerTestCase{
		{"{}", []tokenID{OPEN_CURLY, CLOSE_CURLY}},
		{"][", []tokenID{CLOSE_SQUARE, OPEN_SQUARE}},
		{":true", []tokenID{SEMICOLON, TRUE}}}

	for _, testCase := range testcases {
		scanner := newScanner(testCase.input)
		token := scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[0], token.ID)
		token = scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[1], token.ID)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}

func TestScanner_ShouldStopConcatenatingTokenWhenFindPunctuation(t *testing.T) {

	testcases := []scannerTestCase{
		{"tru,e", []tokenID{INVALID, COMMA, INVALID}},
		{"nul]:", []tokenID{INVALID, CLOSE_SQUARE, SEMICOLON}},
	}

	for _, testCase := range testcases {
		scanner := newScanner(testCase.input)
		token := scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[0], token.ID)
		token = scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[1], token.ID)
		token = scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[2], token.ID)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}

func TestScanner_ShouldReadValidNumbers(t *testing.T) {

	testcases := []scannerTestCase{
		{"132.54,54", []tokenID{NUMBER, COMMA, NUMBER}},
		{"1E+34[-54.5", []tokenID{NUMBER, OPEN_SQUARE, NUMBER}},
	}

	for _, testCase := range testcases {
		scanner := newScanner(testCase.input)
		token := scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[0], token.ID)
		token = scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[1], token.ID)
		token = scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[2], token.ID)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}

func TestScanner_ShouldDetectInValidNumbers(t *testing.T) {

	testcases := []scannerTestCase{
		{"1.32.54,54", []tokenID{INVALID, COMMA, NUMBER}},
		{"1E++34[-54.5E+52", []tokenID{INVALID, OPEN_SQUARE, NUMBER}},
	}

	for _, testCase := range testcases {
		scanner := newScanner(testCase.input)
		token := scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[0], token.ID)
		token = scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[1], token.ID)
		token = scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[2], token.ID)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}

func TestScanner_WillTryToParseAllSymbolsAcceptableForNumbesIntoSingleNumber(t *testing.T) {

	testcases := []scannerTestCase{
		{"1.3+-2.54E54", []tokenID{INVALID}},
		{"1..E+-321.E-0", []tokenID{INVALID}},
	}

	for _, testCase := range testcases {
		scanner := newScanner(testCase.input)
		token := scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[0], token.ID)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}

func TestScanner_ShouldConsiderEverythingInsideQuotesAsString(t *testing.T) {

	testcases := []scannerTestCase{
		{`"This is a string with quotes"`, []tokenID{STRING}},
		{`"1..E+-321.E-0"`, []tokenID{STRING}},
	}

	for _, testCase := range testcases {
		scanner := newScanner(testCase.input)
		token := scanner.NextToken()
		require.Equal(t, testCase.expectedTokens[0], token.ID)
		token = scanner.NextToken()
		require.Equal(t, EOF, token.ID)
	}
}
