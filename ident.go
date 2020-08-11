package builders

import "go/ast"

// Err return err identifier.
func Err() *ast.Ident {
	return ast.NewIdent("err")
}

// Nil return nil identifier.
func Nil() *ast.Ident {
	return ast.NewIdent("nil")
}

// Idents creates array for identifiers.
func Idents(names ...string) []*ast.Ident {
	r := make([]*ast.Ident, len(names))

	for i := range names {
		r[i] = ast.NewIdent(names[i])
	}

	return r
}
