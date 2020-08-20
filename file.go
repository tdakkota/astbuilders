package builders

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"strconv"
)

type FileBuilder struct {
	file *ast.File
	fset *token.FileSet
}

func NewFileBuilder(pkg string) FileBuilder {
	return FileBuilder{
		file: &ast.File{
			Name: ast.NewIdent(pkg),
		},
		fset: token.NewFileSet(),
	}
}

// AddImportPaths adds imports to the file.
func (f FileBuilder) AddImportPaths(paths ...string) FileBuilder {
	for _, path := range paths {
		astutil.AddImport(f.fset, f.file, path)
	}
	return f
}

// AddImports adds imports to the file.
func (f FileBuilder) AddImports(specs ...*ast.ImportSpec) FileBuilder {
	for _, spec := range specs {
		path := spec.Path.Value
		if v, err := strconv.Unquote(spec.Path.Value); err == nil {
			path = v
		}

		if spec.Name == nil {
			astutil.AddImport(f.fset, f.file, path)
		} else {
			astutil.AddNamedImport(f.fset, f.file, spec.Name.Name, path)
		}
	}
	return f
}

// AddDecls adds declarations to the file.
func (f FileBuilder) AddDecls(decls ...ast.Decl) FileBuilder {
	f.file.Decls = append(f.file.Decls, decls...)
	return f
}

// DeclareFunction creates and adds new function declaration to file.
func (f FileBuilder) DeclareFunction(name string, cb func(FunctionBuilder) FunctionBuilder) FileBuilder {
	builder := cb(NewFunctionBuilder(name))
	return f.AddDecls(builder.CompleteAsDecl())
}

// DeclareType creates and adds new type declaration to file.
// If methods parameter is not nil, it is called to add methods.
func (f FileBuilder) DeclareType(name string, typeExpr ast.Expr, methods func(TypeBuilder) TypeBuilder) FileBuilder {
	builder := NewTypeBuilder(name, typeExpr)
	if methods != nil {
		builder = methods(builder)
	}
	return f.AddDecls(builder.CompleteAll()...)
}

// DeclareNewType creates and adds new newtype declaration to file.
func (f FileBuilder) DeclareNewType(name string, typeExpr ast.Expr) FileBuilder {
	return f.DeclareType(name, typeExpr, nil)
}

// DeclareStruct creates and adds new type struct declaration to file.
// If methods parameter is not nil, it is called to add methods.
func (f FileBuilder) DeclareStruct(name string, fields func(StructBuilder) StructBuilder, methods func(TypeBuilder) TypeBuilder) FileBuilder {
	return f.DeclareType(name, fields(NewStructBuilder()).Complete(), methods)
}

// DeclareStruct creates and adds new type struct declaration to file.
// Same as DeclareStruct(name, fields, nil).
func (f FileBuilder) DeclarePlainStruct(name string, fields func(StructBuilder) StructBuilder) FileBuilder {
	return f.DeclareStruct(name, fields, nil)
}

// DeclareInterface creates and adds new type interface declaration to file.
func (f FileBuilder) DeclareInterface(name string, cb func(InterfaceBuilder) InterfaceBuilder) FileBuilder {
	return f.DeclareType(name, cb(NewInterfaceBuilder()).Complete(), nil)
}

// Complete returns built file.
func (f FileBuilder) Complete() *ast.File {
	return f.file
}
