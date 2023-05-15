package jsonquerier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSingleStringField(t *testing.T) {
	payload := `{"Name" : "Bob"}`
	expected := map[string]bool{"Name": true}
	got := Parse(payload)
	require.Equal(t, expected, got)
}
