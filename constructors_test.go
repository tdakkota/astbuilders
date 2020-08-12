package builders

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleParen() {
	par := Paren(ast.NewIdent("parenthesized"))

	printer.Fprint(os.Stdout, token.NewFileSet(), par) // print ast.Node
	// Output: (parenthesized)
}

func ExampleSelectorName() {
	abcd := SelectorName("a", "b", "c", "d")

	printer.Fprint(os.Stdout, token.NewFileSet(), abcd) // print ast.Node
	// Output: a.b.c.d
}

func TestSelectorName(t *testing.T) {
	t.Run("two", func(t *testing.T) {
		sel := SelectorName("a", "b")

		require.Equal(t, sel.X.(*ast.Ident).Name, "a")
		require.Equal(t, sel.Sel.Name, "b")
	})

	t.Run("four", func(t *testing.T) {
		abcd := SelectorName("a", "b", "c", "d")
		abc := abcd.X.(*ast.SelectorExpr)
		ab := abc.X.(*ast.SelectorExpr)

		require.Equal(t, ab.X.(*ast.Ident).Name, "a")
		require.Equal(t, ab.Sel.Name, "b")
		require.Equal(t, abc.Sel.Name, "c")
		require.Equal(t, abcd.Sel.Name, "d")
	})
}
