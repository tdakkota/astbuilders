package builders

import "go/ast"

// Make returns make(typ, length) if capacity is less or equal to 0.
// Otherwise, returns make(typ, length, capacity)
func Make(typ ast.Expr, length, capacity int) *ast.CallExpr {
	var c ast.Expr
	if capacity > 0 {
		c = IntegerLit(capacity)
	}
	return MakeExpr(typ, IntegerLit(length), c)
}

// MakeExpr returns make(typ, length) if capacity is nil
// Otherwise, returns make(typ, length, capacity)
func MakeExpr(typ, length, capacity ast.Expr) *ast.CallExpr {
	args := []ast.Expr{typ, length}
	if capacity != nil {
		args = append(args, capacity)
	}
	return CallName("make", args...)
}

// Len returns len(x).
func Len(x ast.Expr) *ast.CallExpr {
	return CallName("len", x)
}

// Cap returns cap(x).
func Cap(x ast.Expr) *ast.CallExpr {
	return CallName("cap", x)
}

// New returns new(x).
func New(x ast.Expr) *ast.CallExpr {
	return CallName("new", x)
}

// Copy returns copy(dst, src).
func Copy(dst, src ast.Expr) *ast.CallExpr {
	return CallName("copy", dst, src)
}

// Append returns append(slice, elem1, elem2, ..., elemN).
func Append(slice ast.Expr, elem ...ast.Expr) *ast.CallExpr {
	args := append([]ast.Expr{slice}, elem...)
	return CallName("append", args...)
}

// Close returns close(ch).
func Close(ch ast.Expr) *ast.CallExpr {
	return CallName("close", ch)
}
