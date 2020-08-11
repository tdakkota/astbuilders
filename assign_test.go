package builders

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func ExampleSwap() {
	x, y := ast.NewIdent("x"), ast.NewIdent("y")
	swap := Swap(x, y)

	printer.Fprint(os.Stdout, token.NewFileSet(), swap) // print ast.Node
	// Output: x, y = y, x
}
