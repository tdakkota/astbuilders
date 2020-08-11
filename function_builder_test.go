package builders

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunctionBuilder(t *testing.T) {
	a := require.New(t)

	recv := Receiver(ast.NewIdent("r"))(ast.NewIdent("receiver"))
	param := Param(ast.NewIdent("p"))(ast.NewIdent("param"))
	result := Param(ast.NewIdent("res"))(ast.NewIdent("result"))

	f := NewFunctionBuilder("test").
		Recv(recv).
		AddParameters(param).
		AddResults(result).
		Body(func(body StatementBuilder) StatementBuilder {
			return body
		})

	a.Equal(f.recv, recv)
	a.Equal(f.funcType.Params.List[0], param)
	a.Equal(f.funcType.Results.List[0], result)

	call := f.CompleteAsCall()
	a.Equal(call.Fun.(*ast.FuncLit).Type, f.funcType)

	decl := f.CompleteAsDecl()
	a.Equal(decl.Type, f.funcType)
	a.Equal(decl.Name.Name, f.name.Name)
}
