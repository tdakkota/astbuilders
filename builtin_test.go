package builders

import (
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	values := []ast.Expr{
		ast.NewIdent("typ"),
		ast.NewIdent("value"),
		ast.NewIdent("value2"),
	}
	call := Append(values[0], values[1:]...)

	for k, value := range values {
		require.Equal(t, value, call.Args[k])
	}
}

func testBuiltinWithOneParam(t *testing.T, name string, f func(x ast.Expr) *ast.CallExpr) {
	x := ast.NewIdent("ident")

	call := f(x)
	require.Equal(t, name, call.Fun.(*ast.Ident).Name)
	require.Equal(t, x, call.Args[0])
}

func TestLen(t *testing.T) {
	testBuiltinWithOneParam(t, "len", Len)
}

func TestCap(t *testing.T) {
	testBuiltinWithOneParam(t, "cap", Cap)
}

func TestClose(t *testing.T) {
	testBuiltinWithOneParam(t, "close", Close)
}

func TestNew(t *testing.T) {
	testBuiltinWithOneParam(t, "new", New)
}

func TestCopy(t *testing.T) {
	values := []ast.Expr{
		ast.NewIdent("x"),
		ast.NewIdent("y"),
	}
	call := Copy(values[0], values[1])

	for k, value := range values {
		require.Equal(t, value, call.Args[k])
	}
}

func ExampleMake() {
	str := IdentOfKind(types.String)
	slice := SliceOf(str)
	makeCall := Make(slice, 10, 0)

	printer.Fprint(os.Stdout, token.NewFileSet(), makeCall) // print ast.Node
	// Output: make([]string, 10)
}

func ExampleMake_capacity() {
	str := IdentOfKind(types.String)
	slice := SliceOf(str)
	makeCall := Make(slice, 10, 20)

	printer.Fprint(os.Stdout, token.NewFileSet(), makeCall) // print ast.Node
	// Output: make([]string, 10, 20)
}

func TestMakeExpr(t *testing.T) {
	t.Run("with-capacity", func(t *testing.T) {
		values := [...]ast.Expr{
			ast.NewIdent("typ"),
			IntegerLit(10),
			IntegerLit(20),
		}
		call := MakeExpr(values[0], values[1], values[2])

		for k, value := range values {
			require.Equal(t, value, call.Args[k])
		}
	})

	t.Run("with-capacity", func(t *testing.T) {
		values := [...]ast.Expr{
			ast.NewIdent("typ"),
			IntegerLit(10),
		}
		call := MakeExpr(values[0], values[1], nil)

		for k, value := range values {
			require.Equal(t, value, call.Args[k])
		}
	})
}
