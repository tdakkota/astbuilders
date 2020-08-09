package builders

import "go/ast"

func Make(typ ast.Expr, length, capacity int) *ast.CallExpr {
	args := []ast.Expr{typ, IntegerLit(length)}
	if capacity > 0 {
		args = append(args, IntegerLit(capacity))
	}
	return CallName("make", args...)
}

func MakeExpr(typ, length, capacity ast.Expr) *ast.CallExpr {
	args := []ast.Expr{typ, length}
	if capacity != nil {
		args = append(args, capacity)
	}
	return CallName("make", args...)
}

func Len(x ast.Expr) *ast.CallExpr {
	return CallName("len", x)
}

func Cap(x ast.Expr) *ast.CallExpr {
	return CallName("cap", x)
}

func New(x ast.Expr) *ast.CallExpr {
	return CallName("new", x)
}

func Copy(dst, src ast.Expr) *ast.CallExpr {
	return CallName("copy", dst, src)
}

func Append(dst ast.Expr, src ...ast.Expr) *ast.CallExpr {
	args := append([]ast.Expr{dst}, src...)
	return CallName("copy", args...)
}

func Close(ch ast.Expr) *ast.CallExpr {
	return CallName("close", ch)
}
