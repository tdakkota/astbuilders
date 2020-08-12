package builders

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCharLit(t *testing.T) {
	lit := CharLit('a')
	require.Equal(t, "'a'", lit.Value)
	require.Equal(t, token.CHAR, lit.Kind)
}

func TestFloatLit(t *testing.T) {
	lit := FloatLit(1.1)
	require.Equal(t, "1.1", lit.Value)
	require.Equal(t, token.FLOAT, lit.Kind)
}

func TestIntegerLit(t *testing.T) {
	lit := IntegerLit(1)
	require.Equal(t, "1", lit.Value)
	require.Equal(t, token.INT, lit.Kind)
}

func TestStringLit(t *testing.T) {
	lit := StringLit("abc")
	require.Equal(t, `"abc"`, lit.Value)
	require.Equal(t, token.STRING, lit.Kind)
}
