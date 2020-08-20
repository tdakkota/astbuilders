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
	integer := IdentOfKind(types.Int)
	array := ArrayOfSize(integer, 2)
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

func ExampleMapOf() {
	key := IdentOfKind(types.String)
	value := EmptyInterface()
	slice := MapOf(key, value)
	printer.Fprint(os.Stdout, token.NewFileSet(), slice) // print ast.Node
	// Output: map[string]interface {
	// }
}

func ExampleHashSetOf() {
	str := IdentOfKind(types.String)
	slice := HashSetOf(str)
	printer.Fprint(os.Stdout, token.NewFileSet(), slice) // print ast.Node
	// Output: map[string]struct {
	// }
}

func ExampleChanOf() {
	str := EmptyStruct()
	slice := ChanOf(str)
	printer.Fprint(os.Stdout, token.NewFileSet(), slice) // print ast.Node
	// Output: chan struct {
	// }
}

func ExampleSendChanOf() {
	str := EmptyStruct()
	slice := SendChanOf(str)
	printer.Fprint(os.Stdout, token.NewFileSet(), slice) // print ast.Node
	// Output: chan<- struct {
	// }
}

func ExampleRecvChanOf() {
	str := EmptyStruct()
	slice := RecvChanOf(str)
	printer.Fprint(os.Stdout, token.NewFileSet(), slice) // print ast.Node
	// Output: <-chan struct {
	// }
}
