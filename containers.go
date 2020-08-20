package builders

import (
	"go/ast"
)

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

// MapOf returns map[k]v declaration.
func MapOf(k, v ast.Expr) *ast.MapType {
	return &ast.MapType{
		Key:   k,
		Value: v,
	}
}

// HashSetOf returns map[k]struct{} declaration.
func HashSetOf(k ast.Expr) *ast.MapType {
	return MapOf(k, EmptyStruct())
}

// GenericChanOf returns chan elem declaration using given direction.
// Dir is the direction of a channel type is indicated by a bit mask.
// Dir must be ast.SEND, ast.RECV, or ast.SEND | ast.RECV.
func GenericChanOf(dir ast.ChanDir, elem ast.Expr) *ast.ChanType {
	return &ast.ChanType{
		Dir:   dir,
		Value: elem,
	}
}

// ChanOf returns chan elem declaration.
func ChanOf(elem ast.Expr) *ast.ChanType {
	return GenericChanOf(ast.SEND|ast.RECV, elem)
}

// RecvChanOf returns <-chan elem declaration.
func RecvChanOf(elem ast.Expr) *ast.ChanType {
	return GenericChanOf(ast.RECV, elem)
}

// SendChanOf returns chan<- elem declaration.
func SendChanOf(elem ast.Expr) *ast.ChanType {
	return GenericChanOf(ast.SEND, elem)
}
