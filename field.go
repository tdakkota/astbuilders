package builders

import "go/ast"

// Field builds struct field.
func Field(names ...*ast.Ident) func(typ ast.Expr, tag *ast.BasicLit) *ast.Field {
	return func(typ ast.Expr, tag *ast.BasicLit) *ast.Field {
		return &ast.Field{
			Names: names,
			Type:  typ,
			Tag:   tag,
		}
	}
}

// Param builds function parameter.
func Param(names ...*ast.Ident) func(typ ast.Expr) *ast.Field {
	return func(typ ast.Expr) *ast.Field {
		return &ast.Field{
			Names: names,
			Type:  typ,
			Tag:   nil,
		}
	}
}

// Receiver builds function receiver.
func Receiver(name *ast.Ident) func(typ ast.Expr) *ast.Field {
	return Param(name)
}
