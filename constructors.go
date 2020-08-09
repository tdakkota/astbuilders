package builders

import (
	"go/ast"
)

func Call(f ast.Expr, args ...ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  f,
		Args: args,
	}
}

func CallName(name string, args ...ast.Expr) *ast.CallExpr {
	return Call(
		ast.NewIdent(name),
		args...,
	)
}

func CallPackage(pkg, name string, args ...ast.Expr) *ast.CallExpr {
	return Call(&ast.SelectorExpr{
		X:   ast.NewIdent(pkg),
		Sel: ast.NewIdent(name),
	}, args...)
}

func Cast(to, what ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  to,
		Args: []ast.Expr{what},
	}
}

func CastPackage(pkg, name string, what ast.Expr) *ast.CallExpr {
	return Call(&ast.SelectorExpr{
		X:   ast.NewIdent(pkg),
		Sel: ast.NewIdent(name),
	}, what)
}

func FieldName(f *ast.Field) (r *ast.Ident, ok bool) {
	if len(f.Names) < 1 {
		v, ok := f.Type.(*ast.Ident)
		return v, ok
	}

	return f.Names[0], true
}

func Paren(e ast.Expr) *ast.ParenExpr {
	return &ast.ParenExpr{
		X: e,
	}
}

func ArrayOf(e, size ast.Expr) *ast.ArrayType {
	return &ast.ArrayType{
		Len: size,
		Elt: e,
	}
}

func Selector(a ast.Expr, b *ast.Ident, expr ...*ast.Ident) *ast.SelectorExpr {
	if len(expr) == 0 {
		return &ast.SelectorExpr{
			X:   a,
			Sel: b,
		}
	}

	return &ast.SelectorExpr{
		X:   Selector(a, b, expr[:len(expr)-1]...),
		Sel: expr[len(expr)-1],
	}
}

func SelectorName(a, b string, expr ...string) *ast.SelectorExpr {
	exprIdent := make([]*ast.Ident, len(expr))
	for i := range expr {
		exprIdent[i] = ast.NewIdent(expr[i])
	}

	return Selector(ast.NewIdent(a), ast.NewIdent(b), exprIdent...)
}
