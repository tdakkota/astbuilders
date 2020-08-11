package builders_test

import (
	"go/printer"
	"go/token"
	"os"

	"github.com/tdakkota/astbuilders"
)

func Example_ifErrNotNil() {
	// err != nil
	cond := builders.NotEq(builders.Err(), builders.Nil())

	s := builders.NewStatementBuilder()
	s = s.If(nil, cond, func(body builders.StatementBuilder) builders.StatementBuilder {
		return body.Return(builders.Err())
	})

	stmts := s.Complete()
	node := stmts[0]
	printer.Fprint(os.Stdout, token.NewFileSet(), node)
	// Output: if err != nil {
	// 	return err
	// }
}
