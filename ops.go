package builders

import (
	"go/ast"
	"go/token"
)

func DeRef(e ast.Expr) *ast.StarExpr {
	return &ast.StarExpr{
		X: e,
	}
}

func Ref(e ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: token.AND,
		X:  e,
	}
}

func RefFor(ident ast.Expr) *ast.StarExpr {
	return &ast.StarExpr{
		X: ident,
	}
}

func ZeroValue(typ ast.Expr, names ...*ast.Ident) *ast.ValueSpec {
	return &ast.ValueSpec{
		Type:  typ,
		Names: names,
	}
}

type AssignFunc = func(tok token.Token) func(ast.Expr, ...ast.Expr) *ast.AssignStmt

func Assign(lhs1 ast.Expr, lhs ...ast.Expr) AssignFunc {
	return func(tok token.Token) func(ast.Expr, ...ast.Expr) *ast.AssignStmt {
		return func(rhs1 ast.Expr, rhs ...ast.Expr) *ast.AssignStmt {
			return &ast.AssignStmt{
				Lhs: append([]ast.Expr{lhs1}, lhs...),
				Tok: tok,
				Rhs: append([]ast.Expr{rhs1}, rhs...),
			}
		}
	}
}
