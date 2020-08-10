package builders_test

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	"github.com/tdakkota/astbuilders"
)

func Example_ifErrNotNil() {
	errIdent := ast.NewIdent("err")
	nilIdent := ast.NewIdent("nil")
	cond := builders.NotEq(errIdent, nilIdent)

	s := builders.NewStatementBuilder()
	s = s.If(nil, cond, func(body builders.StatementBuilder) builders.StatementBuilder {
		return body.Return(errIdent)
	})

	stmts := s.Complete()
	node := stmts[0]
	printer.Fprint(os.Stdout, token.NewFileSet(), node)
	// Output: if err != nil {
	// 	return err
	// }
}
