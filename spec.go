package builders

import (
	"go/ast"
)

// TypeSpec creates *ast.TypeSpec.
func TypeSpec(name *ast.Ident, typ ast.Expr) *ast.TypeSpec {
	return &ast.TypeSpec{
		Name: name,
		Type: typ,
	}
}

// ValueSpec returns curried function to create *ast.ValueSpec.
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
