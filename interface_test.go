package builders_test

import (
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"os"

	builders "github.com/tdakkota/astbuilders"
)

func ExampleInterfaceBuilder_BuildMethod() {
	i := builders.NewInterfaceBuilder()
	integer := builders.IdentOfKind(types.Int)
	byteType := builders.IdentOfKind(types.Byte)

	i = i.BuildMethod("Write", func(builder builders.FunctionBuilder) builders.FunctionBuilder {
		return builder.
			AddParameters(
				builders.Param(ast.NewIdent("p"))(builders.SliceOf(byteType)),
			).
			AddResults(
				builders.Param(ast.NewIdent("n"))(integer),
				builders.Param(builders.Err())(builders.Error()),
			)
	})

	printer.Fprint(os.Stdout, token.NewFileSet(), i.Complete()) // print ast.Node
	// Output: interface {
	// 	Write(p []uint8) (n int, err error)
	// }
}
