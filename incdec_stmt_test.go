package builders

import (
	"github.com/stretchr/testify/require"
	"go/ast"
	"go/token"
	"testing"
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
