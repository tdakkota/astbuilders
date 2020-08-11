package builders

import "go/ast"

func Index(a, index ast.Expr) *ast.IndexExpr {
	return &ast.IndexExpr{
		X:     a,
		Index: index,
	}
}
