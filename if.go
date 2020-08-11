package builders

import "go/ast"

// If creates if statement.
func If(init ast.Stmt, cond ast.Expr, ifBody BodyFunc) *ast.IfStmt {
	return IfElseStmt(
		init,
		cond,
		ifBody,
		nil,
	)
}

// IfElse creates if-else statement.
func IfElse(init ast.Stmt, cond ast.Expr, ifBody, elseBody BodyFunc) *ast.IfStmt {
	return IfElseStmt(
		init,
		cond,
		ifBody,
		elseBody(NewStatementBuilder()).CompleteAsBlock(),
	)
}

// IfElseStmt creates if-else statement.
// elseStmt should be block or if statement.
func IfElseStmt(init ast.Stmt, cond ast.Expr, ifBody BodyFunc, elseStmt ast.Stmt) *ast.IfStmt {
	return &ast.IfStmt{
		Init: init,
		Cond: cond,
		Body: ifBody(NewStatementBuilder()).CompleteAsBlock(),
		Else: elseStmt,
	}
}
