package builders

import (
	"go/ast"
	"go/token"
)

// AssignFunc represents curried function which creates a assigment statement.
type AssignFunc = func(tok token.Token) func(ast.Expr, ...ast.Expr) *ast.AssignStmt

// Assign returns AssignFunc to build a assigment statement.
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

// Swap returns a, b = b, a statement.
func Swap(a, b ast.Expr) *ast.AssignStmt {
	return Assign(a, b)(token.ASSIGN)(b, a)
}
