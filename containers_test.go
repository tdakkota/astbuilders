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

func TestArrayOf(t *testing.T) {
	i := ast.NewIdent("testType")
	arr := ArrayOf(i, IntegerLit(1))

	require.Equal(t, i, arr.Elt)
	require.Equal(t, "1", arr.Len.(*ast.BasicLit).Value)
}

func ExampleArrayOfSize() {
	int := IdentOfKind(types.Int)
	array := ArrayOfSize(int, 2)
	printer.Fprint(os.Stdout, token.NewFileSet(), array) // print ast.Node
	// Output: [2]int
}

func TestArrayOfSize(t *testing.T) {
	i := ast.NewIdent("testType")
	arr := ArrayOfSize(i, 1)

	require.Equal(t, i, arr.Elt)
	require.Equal(t, "1", arr.Len.(*ast.BasicLit).Value)
}

func ExampleSliceOf() {
	str := IdentOfKind(types.String)
	slice := SliceOf(str)
	printer.Fprint(os.Stdout, token.NewFileSet(), slice) // print ast.Node
	// Output: []string
}

func TestSliceOf(t *testing.T) {
	i := ast.NewIdent("testType")
	arr := SliceOf(i)

	require.Equal(t, i, arr.Elt)
	require.Nil(t, arr.Len)
}
