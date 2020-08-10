package builders

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncStmt(t *testing.T) {
	i := ast.NewIdent("i")
	stmt := IncStmt(i)

	require.Equal(t, i, stmt.X)
	require.Equal(t, token.INC, stmt.Tok)
}

func TestDecStmt(t *testing.T) {
	i := ast.NewIdent("i")
	stmt := DecStmt(i)

	require.Equal(t, i, stmt.X)
	require.Equal(t, token.DEC, stmt.Tok)
}
