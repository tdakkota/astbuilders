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

func ExampleDefine() {
	x := ast.NewIdent("x")
	swap := Define(x)(IntegerLit(0))

	printer.Fprint(os.Stdout, token.NewFileSet(), swap) // print ast.Node
	// Output: x := 0
}
