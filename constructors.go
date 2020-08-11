package builders

import (
	"go/ast"
)

// Paren returns parenthesised expression.
func Paren(e ast.Expr) *ast.ParenExpr {
	return &ast.ParenExpr{
		X: e,
	}
}

// Selector creates selector path from given identifiers.
func Selector(a ast.Expr, b *ast.Ident, expr ...*ast.Ident) *ast.SelectorExpr {
	if len(expr) == 0 {
		return &ast.SelectorExpr{
			X:   a,
			Sel: b,
		}
	}

	return &ast.SelectorExpr{
		X:   Selector(a, b, expr[:len(expr)-1]...),
		Sel: expr[len(expr)-1],
	}
}

// SelectorName creates selector path from given identifiers.
func SelectorName(a, b string, expr ...string) *ast.SelectorExpr {
	idents := Idents(append([]string{a, b}, expr...)...)
	return Selector(idents[0], idents[1], idents[2:]...)
}
