package builders

import (
	"go/ast"
)

func Paren(e ast.Expr) *ast.ParenExpr {
	return &ast.ParenExpr{
		X: e,
	}
}

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

func SelectorName(a, b string, expr ...string) *ast.SelectorExpr {
	exprIdent := make([]*ast.Ident, len(expr))
	for i := range expr {
		exprIdent[i] = ast.NewIdent(expr[i])
	}

	return Selector(ast.NewIdent(a), ast.NewIdent(b), exprIdent...)
}
