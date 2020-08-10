package builders

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func ExampleIndex() {
	call := Index(ast.NewIdent("arr"), IntegerLit(0))
	printer.Fprint(os.Stdout, token.NewFileSet(), call) // print ast.Node
	// Output: arr[0]
}
