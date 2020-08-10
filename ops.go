package builders

import (
	"go/ast"
)

// DeRef returns *e expression.
func DeRef(e ast.Expr) *ast.StarExpr {
	return &ast.StarExpr{
		X: e,
	}
}

// RefFor returns *ident expression.
func RefFor(ident ast.Expr) *ast.StarExpr {
	return &ast.StarExpr{
		X: ident,
	}
}

// TypeAssert returns what.(to) expression.
func TypeAssert(what, to ast.Expr) *ast.TypeAssertExpr {
	return &ast.TypeAssertExpr{
		X:    what,
		Type: to,
	}
}
