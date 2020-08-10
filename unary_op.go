package builders

import (
	"go/ast"
	"go/token"
)

// BinaryOp returns unary expression with given token.
func UnaryOp(e ast.Expr, tok token.Token) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: tok,
		X:  e,
	}
}

// Inc returns e++ expression.
func Inc(e ast.Expr) *ast.UnaryExpr {
	return UnaryOp(e, token.INC)
}

// Dec returns e-- expression.
func Dec(e ast.Expr) *ast.UnaryExpr {
	return UnaryOp(e, token.DEC)
}

// Not returns !e expression.
func Not(e ast.Expr) *ast.UnaryExpr {
	return UnaryOp(e, token.NOT)
}

// BNot returns ^e expression.
func BNot(e ast.Expr) *ast.UnaryExpr {
	return UnaryOp(e, token.XOR)
}

// Recv returns <-e expression.
func Recv(e ast.Expr) *ast.UnaryExpr {
	return UnaryOp(e, token.ARROW)
}

// Ref returns &e expression.
func Ref(e ast.Expr) *ast.UnaryExpr {
	return UnaryOp(e, token.AND)
}
