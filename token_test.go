package jsonquerier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	spelling   string
	expectedID tokenID
}

func TestNewToken_FixedTokens_ShouldBeCategorizedCorrectly(t *testing.T) {
	testcases := []testCase{
		{"{", OPEN_CURLY}, {"}", CLOSE_CURLY},
		{"[", OPEN_SQUARE}, {"]", CLOSE_SQUARE},
		{",", COMMA}, {":", SEMICOLON}, {"true", TRUE},
		{"false", FALSE}, {"null", NULL}}

	for _, testCase := range testcases {
		token := newToken(testCase.spelling)
		require.Equal(t, testCase.expectedID, token.ID)
	}
}

func TestNewToken_NumbersShouldBeCategorizedCorrectly(t *testing.T) {
	numbers := []string{"0", "-5", "5.432", "1.5E+45"}

	for _, number := range numbers {
		token := newToken(number)
		require.Equal(t, NUMBER, token.ID)
	}
}

func TestNewToken_StringssShouldBeCategorizedCorrectly(t *testing.T) {
	strings := []string{`"some string"`, `"another string"`,
		`"\t\n\r\\\/\"\b\f"`, `"+0.52"`}

	for _, str := range strings {
		token := newToken(str)
		require.Equal(t, STRING, token.ID)
	}
}
