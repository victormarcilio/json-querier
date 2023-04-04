package jsonquerier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanner_NextTokenShouldCategorizeSingleCharacterTokens(t *testing.T) {
	singleCharTokens := []string{"{", "}", "[", "]", ",", ":"}

	for _, spelling := range singleCharTokens {
		scanner := newScanner(spelling)
		token := scanner.readNextToken()
		expectedToken := newToken(spelling)
		require.Equal(t, expectedToken, token)
		token = scanner.readNextToken()
		require.Equal(t, EOF, token.ID)
	}
}
