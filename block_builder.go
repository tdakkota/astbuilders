package builders

import (
	"go/ast"
	"go/token"
)

type BodyFunc func(s StatementBuilder) StatementBuilder

type StatementBuilder struct {
	stmts []ast.Stmt
}

func NewStatementBuilder() StatementBuilder {
	return StatementBuilder{}
}

func (builder StatementBuilder) AddStmts(stmts ...ast.Stmt) StatementBuilder {
	builder.stmts = append(builder.stmts, stmts...)
	return builder
}

func (builder StatementBuilder) Expr(expr ast.Expr) StatementBuilder {
	return builder.AddStmts(&ast.ExprStmt{
		X: expr,
	})
}

func (builder StatementBuilder) Decl(decl ast.Decl) StatementBuilder {
	return builder.AddStmts(&ast.DeclStmt{
		Decl: decl,
	})
}

func (builder StatementBuilder) Var(specs ...ast.Spec) StatementBuilder {
	return builder.AddStmts(&ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok:   token.VAR,
			Specs: specs,
		},
	})
}

func (builder StatementBuilder) Const(specs ...ast.Spec) StatementBuilder {
	return builder.AddStmts(&ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok:   token.CONST,
			Specs: specs,
		},
	})
}

func (builder StatementBuilder) If(init ast.Stmt, cond ast.Expr, ifBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.IfStmt{
		Init: init,
		Cond: cond,
		Body: ifBody(NewStatementBuilder()).CompleteAsBlock(),
		Else: nil,
	})
}

func (builder StatementBuilder) IfElse(init ast.Stmt, cond ast.Expr, ifBody, elseBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.IfStmt{
		Init: init,
		Cond: cond,
		Body: ifBody(NewStatementBuilder()).CompleteAsBlock(),
		Else: elseBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

func (builder StatementBuilder) IfElseStmt(init ast.Stmt, cond ast.Expr, ifBody BodyFunc, elseStmt ast.Stmt) StatementBuilder {
	return builder.AddStmts(&ast.IfStmt{
		Init: init,
		Cond: cond,
		Body: ifBody(NewStatementBuilder()).CompleteAsBlock(),
		Else: elseStmt,
	})
}

func (builder StatementBuilder) Block(loopBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.BlockStmt{
		List: loopBody(NewStatementBuilder()).Complete(),
	})
}

func (builder StatementBuilder) For(init ast.Stmt, cond ast.Expr, post ast.Stmt, loopBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.ForStmt{
		Init: init,
		Cond: cond,
		Post: post,
		Body: loopBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

func (builder StatementBuilder) Range(key, value, iter ast.Expr, loopBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.RangeStmt{
		Key:   key,
		Value: value,
		Tok:   token.DEFINE,
		X:     iter,
		Body:  loopBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

func (builder StatementBuilder) Switch(init ast.Stmt, tag ast.Expr, switchBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.SwitchStmt{
		Init: init,
		Tag:  tag,
		Body: switchBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

func (builder StatementBuilder) TypeSwitch(init, assign ast.Stmt, switchBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.TypeSwitchStmt{
		Init:   init,
		Assign: assign,
		Body:   switchBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

func (builder StatementBuilder) CaseExpr(expr ast.Expr, caseBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.CaseClause{
		List: []ast.Expr{expr},
		Body: caseBody(NewStatementBuilder()).Complete(),
	})
}

func (builder StatementBuilder) Case(list []ast.Expr, caseBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.CaseClause{
		List: list,
		Body: caseBody(NewStatementBuilder()).Complete(),
	})
}

func (builder StatementBuilder) Select(selectBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.SelectStmt{
		Body: selectBody(NewStatementBuilder()).CompleteAsBlock(),
	})
}

func (builder StatementBuilder) SelectCase(comm ast.Stmt, caseBody BodyFunc) StatementBuilder {
	return builder.AddStmts(&ast.CommClause{
		Comm: comm,
		Body: caseBody(NewStatementBuilder()).Complete(),
	})
}

func (builder StatementBuilder) Branch(token token.Token) StatementBuilder {
	return builder.AddStmts(&ast.BranchStmt{
		Tok: token,
	})
}

func (builder StatementBuilder) BranchLabel(token token.Token, label *ast.Ident) StatementBuilder {
	return builder.AddStmts(&ast.BranchStmt{
		Tok:   token,
		Label: label,
	})
}

func (builder StatementBuilder) Go(call *ast.CallExpr) StatementBuilder {
	return builder.AddStmts(&ast.GoStmt{
		Call: call,
	})
}

func (builder StatementBuilder) Defer(call *ast.CallExpr) StatementBuilder {
	return builder.AddStmts(&ast.DeferStmt{
		Call: call,
	})
}

func (builder StatementBuilder) Define(lhs1 ast.Expr, lhs ...ast.Expr) func(ast.Expr, ...ast.Expr) StatementBuilder {
	return builder.Assign(lhs1, lhs...)(token.DEFINE)
}

func (builder StatementBuilder) Assign(lhs1 ast.Expr, lhs ...ast.Expr) func(tok token.Token) func(ast.Expr, ...ast.Expr) StatementBuilder {
	return func(tok token.Token) func(ast.Expr, ...ast.Expr) StatementBuilder {
		return func(rhs1 ast.Expr, rhs ...ast.Expr) StatementBuilder {
			return builder.AddStmts(&ast.AssignStmt{
				Lhs: append([]ast.Expr{lhs1}, lhs...),
				Tok: tok,
				Rhs: append([]ast.Expr{rhs1}, rhs...),
			})
		}
	}
}

func (builder StatementBuilder) Return(results ...ast.Expr) StatementBuilder {
	return builder.AddStmts(&ast.ReturnStmt{
		Results: results,
	})
}

func (builder StatementBuilder) Add(add StatementBuilder) StatementBuilder {
	return builder.AddStmts(add.Complete()...)
}

func (builder StatementBuilder) Complete() []ast.Stmt {
	return builder.stmts
}

func (builder StatementBuilder) CompleteAsBlock() *ast.BlockStmt {
	return &ast.BlockStmt{
		List: builder.Complete(),
	}
}
