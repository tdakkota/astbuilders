package builders

import "go/ast"

// ArrayOf returns [size]elem declaration.
func ArrayOf(elem, size ast.Expr) *ast.ArrayType {
	return &ast.ArrayType{
		Len: size,
		Elt: elem,
	}
}

// ArrayOfSize returns [size]elem declaration.
func ArrayOfSize(e ast.Expr, size int) *ast.ArrayType {
	return &ast.ArrayType{
		Len: IntegerLit(size),
		Elt: e,
	}
}

// SliceOf returns []elem declaration.
func SliceOf(elem ast.Expr) *ast.ArrayType {
	return &ast.ArrayType{
		Len: nil,
		Elt: elem,
	}
}
