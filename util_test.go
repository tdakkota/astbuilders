package builders

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFieldName(t *testing.T) {
	t.Run("with-names", func(t *testing.T) {
		field := &ast.Field{
			Names: []*ast.Ident{ast.NewIdent("name")},
			Type:  ast.NewIdent("typ"),
		}

		r, ok := FieldName(field)
		require.True(t, ok)
		require.Equal(t, field.Names[0], r)
	})

	t.Run("only-type", func(t *testing.T) {
		typ := ast.NewIdent("typ")
		field := &ast.Field{
			Type: typ,
		}

		r, ok := FieldName(field)
		require.True(t, ok)
		require.Equal(t, typ, r)
	})
}
