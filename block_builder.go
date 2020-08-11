package builders

import (
	"go/ast"
	"go/token"
)

type BodyFunc func(s StatementBuilder) StatementBuilder

type StatementBuilder struct {
	stmts []ast.Stmt
}

// NewStatementBuilder creates new StatementBuilder.
func NewStatementBuilder() StatementBuilder {
	return StatementBuilder{}
}

// AddStmts adds statements to block.
func (builder StatementBuilder) AddStmts(stmts ...ast.Stmt) StatementBuilder {
	builder.stmts = append(builder.stmts, stmts...)
	return builder
}

// Expr adds expression to block.
func (builder StatementBuilder) Expr(expr ast.Expr) StatementBuilder {
	return builder.AddStmts(&ast.ExprStmt{
		X: expr,
	})
}

// Decl adds declaration to block.
func (builder StatementBuilder) Decl(decl ast.Decl) StatementBuilder {
	return builder.AddStmts(&ast.DeclStmt{
		Decl: decl,
	})
}

// GenDecl adds declarations to block.
// tok should be token.VAR, token.CONST, token.TYPE.
// Also token.IMPORT can be used outside function blocks.
func (builder StatementBuilder) GenDecl(tok token.Token, specs ...ast.Spec) StatementBuilder {
	return builder.Decl(&ast.GenDecl{
		Tok:   tok,
		Specs: specs,
	})
}

// Var adds variable declarations to block.
func (builder StatementBuilder) Var(specs ...ast.Spec) StatementBuilder {
	return builder.GenDecl(
		token.VAR,
		specs...,
	)
}

// Const adds constant declarations to block.
func (builder StatementBuilder) Const(specs ...ast.Spec) StatementBuilder {
	return builder.GenDecl(
		token.CONST,
		specs...,
	)
}

// If adds if statement.
func (builder StatementBuilder) If(init ast.Stmt, cond ast.Expr, ifBody BodyFunc) StatementBuilder {
	return builder.AddStmts(If(init, cond, ifBody))
}

// IfElse adds if-else statement.
func (builder StatementBuilder) IfElse(init ast.Stmt, cond ast.Expr, ifBody, elseBody BodyFunc) StatementBuilder {
	return builder.AddStmts(IfElse(init, cond, ifBody, elseBody))
}

// IfElseStmt adds if-else statement.
// elseStmt should be block or if statement.
func (builder StatementBuilder) IfElseStmt(init ast.Stmt, cond ast.Expr, ifBody BodyFunc, elseStmt ast.Stmt) StatementBuilder {
	return builder.AddStmts(IfElseStmt(init, cond, ifBody, elseStmt))
}

// Block adds block statement.
func (builder StatementBuilder) Block(loopBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.BlockStmt{
		List: loopBody(NewStatementBuilder()).Complete(),
	})
}

// For adds for statement.
func (builder StatementBuilder) For(init ast.Stmt, cond ast.Expr, post ast.Stmt, loopBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.ForStmt{
		Init: init,
		Cond: cond,
		Post: post,
		Body: loopBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

// Range adds range statement.
func (builder StatementBuilder) Range(key, value, iter ast.Expr, loopBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.RangeStmt{
		Key:   key,
		Value: value,
		Tok:   token.DEFINE,
		X:     iter,
		Body:  loopBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

// Switch adds switch statement.
func (builder StatementBuilder) Switch(init ast.Stmt, tag ast.Expr, switchBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.SwitchStmt{
		Init: init,
		Tag:  tag,
		Body: switchBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

// TypeSwitch adds type-switch statement.
func (builder StatementBuilder) TypeSwitch(init, assign ast.Stmt, switchBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.TypeSwitchStmt{
		Init:   init,
		Assign: assign,
		Body:   switchBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

// CaseExpr adds case statement.
func (builder StatementBuilder) CaseExpr(expr ast.Expr, caseBody BodyFunc) StatementBuilder {
	return builder.Case([]ast.Expr{expr}, caseBody)
}

// Case adds case statement.
func (builder StatementBuilder) Case(list []ast.Expr, caseBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.CaseClause{
		List: list,
		Body: caseBody(NewStatementBuilder()).Complete(),
	})
}

// Select adds select statement.
func (builder StatementBuilder) Select(selectBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.SelectStmt{
		Body: selectBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

// SelectCase adds case of select statement.
func (builder StatementBuilder) SelectCase(comm ast.Stmt, caseBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.CommClause{
		Comm: comm,
		Body: caseBody(NewStatementBuilder()).Complete(),
	})
}

// Branch adds branch statement.
// token should be token.BREAK, token.CONTINUE, token.GOTO or token.FALLTHROUGH
func (builder StatementBuilder) Branch(tok token.Token) StatementBuilder {
	return builder.BranchLabel(tok, nil)
}

// BranchLabel adds labeled branch statement.
func (builder StatementBuilder) BranchLabel(tok token.Token, label *ast.Ident) StatementBuilder {
	return builder.AddStmts(&ast.BranchStmt{
		Tok:   tok,
		Label: label,
	})
}

// Go adds go statement.
func (builder StatementBuilder) Go(call *ast.CallExpr) StatementBuilder {
	return builder.AddStmts(&ast.GoStmt{
		Call: call,
	})
}

// Defer adds defer statement.
func (builder StatementBuilder) Defer(call *ast.CallExpr) StatementBuilder {
	return builder.AddStmts(&ast.DeferStmt{
		Call: call,
	})
}

// Define adds define statement.
func (builder StatementBuilder) Define(lhs1 ast.Expr, lhs ...ast.Expr) func(ast.Expr, ...ast.Expr) StatementBuilder {
	return builder.Assign(lhs1, lhs...)(token.DEFINE)
}

// Assign adds assign statement.
func (builder StatementBuilder) Assign(lhs1 ast.Expr, lhs ...ast.Expr) func(tok token.Token) func(ast.Expr, ...ast.Expr) StatementBuilder {
	return func(tok token.Token) func(ast.Expr, ...ast.Expr) StatementBuilder {
		return func(rhs1 ast.Expr, rhs ...ast.Expr) StatementBuilder {
			return builder.AddStmts(Assign(lhs1, lhs...)(tok)(rhs1, rhs...))
		}
	}
}

// Return adds return statement.
func (builder StatementBuilder) Return(results ...ast.Expr) StatementBuilder {
	return builder.AddStmts(&ast.ReturnStmt{
		Results: results,
	})
}

// Add adds statements from another statement builder.
func (builder StatementBuilder) Add(add StatementBuilder) StatementBuilder {
	return builder.AddStmts(add.Complete()...)
}

// Complete returns all added statements.
func (builder StatementBuilder) Complete() []ast.Stmt {
	return builder.stmts
}

// Complete returns all added statements as block statement.
func (builder StatementBuilder) CompleteAsBlock() *ast.BlockStmt {
	return &ast.BlockStmt{
		List: builder.Complete(),
	}
}
