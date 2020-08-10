package builders

import (
	"go/ast"
	"go/types"
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

func TestTypeAssert(t *testing.T) {
	x := ast.NewIdent("x")
	i := IdentOfKind(types.Int)
	assert := TypeAssert(x, i)

	require.Equal(t, x, assert.X)
	require.Equal(t, i, assert.Type)
}
