package builders

import (
	"go/ast"
	"go/token"
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
	i := ast.NewIdent("i")
	zero := IntegerLit(0)
	init := Define(i)(zero)
	cond := Less(i, zero)

	t.Run("if", func(t *testing.T) {
		a := require.New(t)
		s := NewStatementBuilder()

		s = s.If(init, cond, func(body StatementBuilder) StatementBuilder {
			return body.Return(i)
		})

		stmt := s.stmts[0].(*ast.IfStmt)
		a.Equal(init, stmt.Init)
		a.Equal(cond, stmt.Cond)
		a.Len(stmt.Body.List, 1)
		a.Nil(stmt.Else)
	})

	t.Run("else", func(t *testing.T) {
		a := require.New(t)
		s := NewStatementBuilder()

		s = s.IfElse(init, cond, func(body StatementBuilder) StatementBuilder {
			return body.Return(i)
		}, func(body StatementBuilder) StatementBuilder {
			return body.Return(i)
		})

		stmt := s.stmts[0].(*ast.IfStmt)
		a.Equal(init, stmt.Init)
		a.Equal(cond, stmt.Cond)
		a.Len(stmt.Body.List, 1)
		a.Len(stmt.Else.(*ast.BlockStmt).List, 1)
	})
}

func TestDefine(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	x, y := ast.NewIdent("x"), ast.NewIdent("y")
	s = s.Define(x)(y)

	stmt := s.stmts[0].(*ast.AssignStmt)
	a.Equal(x, stmt.Lhs[0])
	a.Equal(y, stmt.Rhs[0])
}

func TestBlock(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	s = s.Block(func(body StatementBuilder) StatementBuilder {
		return body.Branch(token.CONTINUE)
	})

	stmt := s.stmts[0].(*ast.BlockStmt)
	a.Len(stmt.List, 1)
	a.Equal(token.CONTINUE, stmt.List[0].(*ast.BranchStmt).Tok)
}

func TestRange(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	k := ast.NewIdent("k")
	iter := ast.NewIdent("iter")
	s = s.Range(k, nil, iter, func(s StatementBuilder) StatementBuilder {
		return s
	})

	stmt := s.stmts[0].(*ast.RangeStmt)
	a.Equal(k, stmt.Key)
	a.Nil(stmt.Value)
	a.Equal(iter, stmt.X)
	a.Len(stmt.Body.List, 0)
}

func TestFor(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	i := ast.NewIdent("i")
	zero := IntegerLit(0)
	init := Define(i)(zero)
	cond := Less(i, zero)
	post := IncStmt(i)

	s = s.For(init, cond, post, func(loop StatementBuilder) StatementBuilder {
		return loop
	})

	stmt := s.stmts[0].(*ast.ForStmt)
	a.Equal(init, stmt.Init)
	a.Equal(cond, stmt.Cond)
	a.Equal(post, stmt.Post)
}

func TestSwitch(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	i := ast.NewIdent("i")
	zero := IntegerLit(0)
	init := Define(i)(zero)
	cond := Less(i, zero)

	s = s.Switch(init, cond, func(body StatementBuilder) StatementBuilder {
		body = body.CaseExpr(i, func(cse StatementBuilder) StatementBuilder {
			return cse
		})
		return body
	})

	stmt := s.stmts[0].(*ast.SwitchStmt)
	a.Equal(init, stmt.Init)
	a.Equal(cond, stmt.Tag)
	a.Equal(i, stmt.Body.List[0].(*ast.CaseClause).List[0])
}

func TestTypeSwitch(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	i := ast.NewIdent("i")
	zero := IntegerLit(0)
	init := Define(i)(zero)
	assign := &ast.ExprStmt{
		X: ast.NewIdent("int"),
	}

	s = s.TypeSwitch(init, assign, func(body StatementBuilder) StatementBuilder {
		body = body.CaseExpr(i, func(cse StatementBuilder) StatementBuilder {
			return cse
		})
		return body
	})

	stmt := s.stmts[0].(*ast.TypeSwitchStmt)
	a.Equal(init, stmt.Init)
	a.Equal(assign, stmt.Assign)
	a.Equal(i, stmt.Body.List[0].(*ast.CaseClause).List[0])
}

func TestSelect(t *testing.T) {
	a := require.New(t)
	s := NewStatementBuilder()

	ch := &ast.ExprStmt{
		X: Recv(ast.NewIdent("ch")),
	}

	s = s.Select(func(body StatementBuilder) StatementBuilder {
		body = body.SelectCase(ch, func(cse StatementBuilder) StatementBuilder {
			return cse
		})
		return body
	})

	stmt := s.stmts[0].(*ast.SelectStmt)
	a.Equal(ch, stmt.Body.List[0].(*ast.CommClause).Comm)
}
