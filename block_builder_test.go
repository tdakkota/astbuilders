package builders

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComplete(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	expr := ast.NewIdent("i")
	s = s.Expr(expr)

	a.Equal(s.stmts[0].(*ast.ExprStmt).X, expr)
	a.Len(s.stmts, 1)
	a.Equal(s.stmts, s.Complete())
	a.Equal(s.stmts, s.CompleteAsBlock().List)

	s = s.Add(s)
	a.Len(s.stmts, 2)
	a.Equal(s.stmts, s.Complete())
	a.Equal(s.stmts, s.CompleteAsBlock().List)
}

func TestGenDecl(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	ident := ast.NewIdent("i")
	spec := ValueSpec(ident)(nil)(IntegerLit(0))

	s = s.Var(spec)
	stmts := s.Complete()
	a.Len(stmts, 1)
	a.Equal(stmts[0].(*ast.DeclStmt).Decl.(*ast.GenDecl).Specs[0], spec)

	s = s.Const(spec)
	stmts = s.Complete()
	a.Len(stmts, 2)
	a.Equal(stmts[1].(*ast.DeclStmt).Decl.(*ast.GenDecl).Specs[0], spec)
}

func TestCall(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	call := CallName("f")

	s = s.Defer(call)
	stmts := s.Complete()
	a.Len(stmts, 1)
	a.Equal(stmts[0].(*ast.DeferStmt).Call, call)

	s = s.Go(call)
	stmts = s.Complete()
	a.Len(stmts, 2)
	a.Equal(stmts[1].(*ast.GoStmt).Call, call)
}

func TestIf(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	i := ast.NewIdent("i")
	zero := IntegerLit(0)
	init := Define(i)(zero)
	cond := Less(i, zero)
	s = s.If(init, cond, func(body StatementBuilder) StatementBuilder {
		return body.Return(i)
	})

	stmt := s.stmts[0].(*ast.IfStmt)
	a.Equal(init, stmt.Init)
	a.Equal(cond, stmt.Cond)
	a.Len(stmt.Body.List, 1)
	a.Nil(stmt.Else)
}
