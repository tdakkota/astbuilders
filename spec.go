package builders

import (
	"go/ast"
)

func ValueSpec(name *ast.Ident, names ...*ast.Ident) func(typ ast.Expr) func(...ast.Expr) *ast.ValueSpec {
	return func(typ ast.Expr) func(...ast.Expr) *ast.ValueSpec {
		return func(exprs ...ast.Expr) *ast.ValueSpec {
			return &ast.ValueSpec{
				Names:  append([]*ast.Ident{name}, names...),
				Type:   typ,
				Values: exprs,
			}
		}
	}
}
