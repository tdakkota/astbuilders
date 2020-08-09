package builders

import (
	"go/ast"
	"go/token"
)

func Inc(e ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: token.INC,
		X:  e,
	}
}

func Dec(e ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: token.DEC,
		X:  e,
	}
}

func Not(e ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: token.NOT,
		X:  e,
	}
}
