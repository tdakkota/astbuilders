package builders

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeRef(t *testing.T) {
	x := ast.NewIdent("x")
	star := DeRef(x)

	require.Equal(t, x, star.X)
}

func TestRefFor(t *testing.T) {
	x := ast.NewIdent("x")
	star := RefFor(x)

	require.Equal(t, x, star.X)
}
