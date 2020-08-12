package builders

import (
	"go/ast"
	"go/token"
)

// FieldName gets one name of field.
func FieldName(f *ast.Field) (r *ast.Ident, ok bool) {
	names, ok := FieldNames(f)

	return names[0], ok
}

// FieldName gets all names of field.
func FieldNames(f *ast.Field) (r []*ast.Ident, ok bool) {
	if len(f.Names) < 1 {
		v, ok := f.Type.(*ast.Ident)
		return []*ast.Ident{v}, ok
	}

	return f.Names, true
}

// Imports creates import declaration from imports specs.
func Imports(imports ...*ast.ImportSpec) *ast.GenDecl {
	specs := make([]ast.Spec, len(imports))
	for i := range imports {
		specs[i] = imports[i]
	}

	return &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: specs,
	}
}
