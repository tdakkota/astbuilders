package builders

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryOps(t *testing.T) {
	expect := map[token.Token]func(x ast.Expr, y ast.Expr) *ast.BinaryExpr{
		token.ADD: Add, // +
		token.SUB: Sub, // -
		token.MUL: Mul, // *
		token.QUO: Div, // /
		token.REM: Rem, // %

		token.AND:     BAnd,       // &
		token.OR:      BOr,        // |
		token.XOR:     Xor,        // ^
		token.SHL:     ShiftLeft,  // <<
		token.SHR:     ShiftRight, // >>
		token.AND_NOT: AddNot,     // &^

		token.LAND: And, // &&
		token.LOR:  Or,  // ||

		token.EQL: Eq,      // ==
		token.LSS: Less,    // <
		token.GTR: Greater, // >

		token.NEQ: NotEq,       // !=
		token.LEQ: LessOrEq,    // <=
		token.GEQ: GreaterOrEq, // >=

		token.ARROW: Send, // <-
	}

	r := require.New(t)
	x, y := ast.NewIdent("x"), ast.NewIdent("y")
	for tok, f := range expect {
		expr := f(x, y)
		r.Equal(tok, expr.Op)
		r.Equal(x, expr.X)
		r.Equal(y, expr.Y)
	}
}
