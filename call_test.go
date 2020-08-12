package builders

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleCall() {
	call := Call(ast.NewIdent("f"), IntegerLit(0))
	printer.Fprint(os.Stdout, token.NewFileSet(), call) // print ast.Node
	// Output: f(0)
}

func TestCallName(t *testing.T) {
	expr := CallName("testCall", IntegerLit(1))

	require.Equal(t, "testCall", expr.Fun.(*ast.Ident).Name)
	require.Equal(t, "1", expr.Args[0].(*ast.BasicLit).Value)
}

func ExampleCallPackage() {
	call := CallPackage("fmt", "Println", StringLit("Hello, World!"))
	printer.Fprint(os.Stdout, token.NewFileSet(), call) // print ast.Node
	// Output: fmt.Println("Hello, World!")
}

func TestCallPackage(t *testing.T) {
	expr := CallPackage("testpkg", "testCall", IntegerLit(1))
	sel := expr.Fun.(*ast.SelectorExpr)

	require.Equal(t, "testCall", sel.Sel.Name)
	require.Equal(t, "testpkg", sel.X.(*ast.Ident).Name)
	require.Equal(t, "1", expr.Args[0].(*ast.BasicLit).Value)
}

func TestCastPackage(t *testing.T) {
	expr := CastPackage("testpkg", "testCast", IntegerLit(1))
	sel := expr.Fun.(*ast.SelectorExpr)

	require.Equal(t, "testCast", sel.Sel.Name)
	require.Equal(t, "testpkg", sel.X.(*ast.Ident).Name)
	require.Equal(t, "1", expr.Args[0].(*ast.BasicLit).Value)
}
