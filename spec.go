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

// ImportSpec creates *ast.ImportSpec.
func ImportSpec(name *ast.Ident, path *ast.BasicLit) *ast.ImportSpec {
	return &ast.ImportSpec{
		Name: name,
		Path: path,
	}
}

// Import create unnamed import spec using path.
func Import(path string) *ast.ImportSpec {
	return ImportSpec(nil, StringLit(path))
}

// Import create unnamed import spec using path.
func NamedImport(name, path string) *ast.ImportSpec {
	return ImportSpec(ast.NewIdent(name), StringLit(path))
}
