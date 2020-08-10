package builders

import (
	"go/ast"
	"go/token"
)

// IncDecStmt returns increment/decrement statement.
func IncDecStmt(x ast.Expr, tok token.Token) *ast.IncDecStmt {
	return &ast.IncDecStmt{
		X:   x,
		Tok: tok,
	}
}

// IncStmt returns increment statement.
func IncStmt(x ast.Expr) *ast.IncDecStmt {
	return IncDecStmt(x, token.INC)
}

// DecStmt returns decrement statement.
func DecStmt(x ast.Expr) *ast.IncDecStmt {
	return IncDecStmt(x, token.DEC)
}
