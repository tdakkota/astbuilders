package builders

import (
	"github.com/stretchr/testify/require"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestFileBuilder_DeclareType(t *testing.T) {
	a := require.New(t)
	f := NewFileBuilder("main")

	a.Equal("main", f.file.Name.Name)
	integer := IdentOfKind(types.Int)

	var typeDecl *ast.GenDecl
	var methodDecl *ast.FuncDecl
	f = f.DeclareType("typ", SliceOf(integer), func(typ TypeBuilder) TypeBuilder {
		recv := ast.NewIdent("r")

		typ = typ.AddMethod("Add", recv, func(add FunctionBuilder) FunctionBuilder {
			b := ast.NewIdent("b")
			param := Param(b)(typ.Ident())
			result := Param()(typ.Ident())

			add = add.SetType(FuncType(param)(result)).Body(func(body StatementBuilder) StatementBuilder {
				return body.Return(Add(recv, b))
			})
			methodDecl = add.CompleteAsDecl()

			return add
		})
		typeDecl = typ.CompleteAsDecl(token.TYPE)

		return typ
	})

	f = f.DeclareNewType("newtyp", ast.NewIdent("typ"))

	file := f.Complete()
	a.Len(file.Decls, 3)
	a.Equal(typeDecl, file.Decls[0])
	a.Equal(methodDecl, file.Decls[1])

	newtyp := file.Decls[2].(*ast.GenDecl)
	a.Equal(ast.NewIdent("typ"), newtyp.Specs[0].(*ast.TypeSpec).Type)
}

func TestFileBuilder_DeclareStruct(t *testing.T) {
	a := require.New(t)
	f := NewFileBuilder("main")
	integer := IdentOfKind(types.Int)

	var s *ast.StructType
	f = f.DeclarePlainStruct("typ", func(builder StructBuilder) StructBuilder {
		builder = builder.AddFields(Field(Idents("integer")...)(integer, nil))
		s = builder.Complete()
		return builder
	})

	file := f.Complete()
	a.Len(file.Decls, 1)
	a.Equal(s, file.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type)
}

func TestFileBuilder_DeclareInterface(t *testing.T) {
	a := require.New(t)
	f := NewFileBuilder("main")

	var s *ast.InterfaceType
	f = f.DeclareInterface("Stringer", func(builder InterfaceBuilder) InterfaceBuilder {
		builder = builder.AddMethod("String", FuncType()())
		s = builder.Complete()
		return builder
	})

	file := f.Complete()
	a.Len(file.Decls, 1)
	a.Equal(s, file.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type)
}

func TestFileBuilder_DeclareFunction(t *testing.T) {
	a := require.New(t)
	f := NewFileBuilder("main")

	var s *ast.FuncType
	f = f.DeclareFunction("String", func(builder FunctionBuilder) FunctionBuilder {
		s = FuncType()()
		return builder.SetType(s)
	})

	file := f.Complete()
	a.Len(file.Decls, 1)
	a.Equal(s, file.Decls[0].(*ast.FuncDecl).Type)
}

func TestFileBuilder_Imports(t *testing.T) {
	a := require.New(t)
	f := NewFileBuilder("main")

	f = f.AddImportPaths("github.com/tdakkota/astbuilders")
	f = f.AddImports(Import("github.com/tdakkota/gomacro"))
	f = f.AddImports(NamedImport("macros", "github.com/tdakkota/gomacro"))
	file := f.Complete()

	a.Len(file.Decls, 1)
	specs := file.Decls[0].(*ast.GenDecl).Specs
	a.Len(specs, 3)

	a.Equal(`"github.com/tdakkota/astbuilders"`, specs[0].(*ast.ImportSpec).Path.Value)
	a.Equal("macros", specs[1].(*ast.ImportSpec).Name.Name)
	a.Equal(`"github.com/tdakkota/gomacro"`, specs[1].(*ast.ImportSpec).Path.Value)
	a.Equal(`"github.com/tdakkota/gomacro"`, specs[2].(*ast.ImportSpec).Path.Value)
}
