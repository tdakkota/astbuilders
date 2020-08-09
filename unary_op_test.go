package builders

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnaryOps(t *testing.T) {
	expect := map[token.Token]func(e ast.Expr) *ast.UnaryExpr{
		token.INC:   Inc,  // ++
		token.DEC:   Dec,  // --
		token.NOT:   Not,  // !
		token.XOR:   BNot, // ^
		token.ARROW: Recv, // <-
		token.AND:   Ref,  // &
	}

	r := require.New(t)
	e := ast.NewIdent("e")
	for tok, f := range expect {
		expr := f(e)
		r.Equal(tok, expr.Op)
		r.Equal(e, expr.X)
	}
}
