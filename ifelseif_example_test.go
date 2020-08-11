package builders_test

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	"github.com/tdakkota/astbuilders"
)

func returnString(lit string) builders.BodyFunc {
	return func(body builders.StatementBuilder) builders.StatementBuilder {
		return body.Return(builders.StringLit(lit))
	}
}

func Example_ifElseIf() {
	a := ast.NewIdent("a")
	b := ast.NewIdent("b")

	s := builders.NewStatementBuilder()
	lessCase := builders.IfElseStmt(nil, builders.Less(a, b), returnString("less"), nil)
	greaterCase := builders.IfElseStmt(nil, builders.Greater(a, b), returnString("greater"), lessCase)
	s = s.IfElseStmt(nil, builders.Eq(a, b), returnString("equal"), greaterCase)

	stmts := s.Complete()
	node := stmts[0]
	printer.Fprint(os.Stdout, token.NewFileSet(), node)
	// Output: if a == b {
	// 	return "equal"
	// } else if a > b {
	// 	return "greater"
	// } else if a < b {
	// 	return "less"
	// }
}
