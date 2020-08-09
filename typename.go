package builders

import (
	"go/ast"
	"go/types"
)

func IdentOfKind(kind types.BasicKind) *ast.Ident {
	return ast.NewIdent(types.Typ[kind].Name())
}
