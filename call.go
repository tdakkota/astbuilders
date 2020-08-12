package builders

import "go/ast"

// Call returns f(args).
func Call(f ast.Expr, args ...ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  f,
		Args: args,
	}
}

// CallName returns name(args).
// Fast path for Call.
func CallName(name string, args ...ast.Expr) *ast.CallExpr {
	return Call(
		ast.NewIdent(name),
		args...,
	)
}

// CallPackage returns pkg.name(args).
// Fast path for Call.
func CallPackage(pkg, name string, args ...ast.Expr) *ast.CallExpr {
	return Call(&ast.SelectorExpr{
		X:   ast.NewIdent(pkg),
		Sel: ast.NewIdent(name),
	}, args...)
}

// Cast returns to(what).
func Cast(to, what ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  to,
		Args: []ast.Expr{what},
	}
}

// CastPackage returns pkg.name(what).
// Fast path for Cast.
func CastPackage(pkg, name string, what ast.Expr) *ast.CallExpr {
	return Cast(&ast.SelectorExpr{
		X:   ast.NewIdent(pkg),
		Sel: ast.NewIdent(name),
	}, what)
}
