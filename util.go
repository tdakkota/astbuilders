package builders

import "go/ast"

func FieldName(f *ast.Field) (r *ast.Ident, ok bool) {
	names, ok := FieldNames(f)

	return names[0], ok
}

func FieldNames(f *ast.Field) (r []*ast.Ident, ok bool) {
	if len(f.Names) < 1 {
		v, ok := f.Type.(*ast.Ident)
		return []*ast.Ident{v}, ok
	}

	return f.Names, true
}
