package builders

import (
	"go/ast"
	"go/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypeBuilder(t *testing.T) {
	integer := IdentOfKind(types.Int)
	b := NewTypeBuilder("MyInt", integer)
	b = b.AddMethod("testMethod", ast.NewIdent("r"), func(builder FunctionBuilder) FunctionBuilder {
		param := Param(ast.NewIdent("a"))(b.Ident())
		return builder.AddParameters(param)
	})

	a := require.New(t)
	decls := b.CompleteAll()
	a.Len(decls, 2)

	typDecl := decls[0].(*ast.GenDecl)
	spec := typDecl.Specs[0].(*ast.TypeSpec)
	a.Equal("MyInt", spec.Name.Name)
	a.Equal(integer, spec.Type)

	method := decls[1].(*ast.FuncDecl)
	a.Equal("testMethod", method.Name.Name)
	// check receiver
	a.Equal("r", method.Recv.List[0].Names[0].Name)
	a.Equal(b.Ident(), method.Recv.List[0].Type)
	// check parameter
	typ := method.Type
	a.Equal("a", typ.Params.List[0].Names[0].Name)
	a.Equal(b.Ident(), typ.Params.List[0].Type)
}
