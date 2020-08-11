package builders

import (
	"go/ast"
	"go/token"
)

// TypeBuilder is type declaration builder.
type TypeBuilder struct {
	name     string
	typeExpr ast.Expr
	methods  []*ast.FuncDecl
}

// NewTypeBuilder creates TypeBuilder.
func NewTypeBuilder(name string, typeExpr ast.Expr) TypeBuilder {
	return TypeBuilder{name: name, typeExpr: typeExpr}
}

// Ident returns identifier of building type.
func (s TypeBuilder) Ident() *ast.Ident {
	return ast.NewIdent(s.name)
}

// AddMethod creates and adds an associated method.
func (s TypeBuilder) AddMethod(
	methodName string,
	receiver *ast.Ident,
	f func(FunctionBuilder) FunctionBuilder,
) TypeBuilder {
	builder := NewFunctionBuilder(methodName)
	builder = builder.Recv(Receiver(receiver)(s.Ident()))

	return s.AddMethodDecls(f(builder).CompleteAsDecl())
}

// AddMethodDecls adds associated methods.
func (s TypeBuilder) AddMethodDecls(decls ...*ast.FuncDecl) TypeBuilder {
	s.methods = append(s.methods, decls...)
	return s
}

// CompleteMethods returns associated methods.
func (s TypeBuilder) CompleteMethods() []*ast.FuncDecl {
	return s.methods
}

// CompleteAsSpec returns a type declaration with built struct.
// i.e. name struct{/* fields */}
func (s TypeBuilder) CompleteAsSpec() *ast.TypeSpec {
	return TypeSpec(s.Ident(), s.typeExpr)
}

// CompleteAsDecl returns a type declaration with built struct
// e.g type name struct{} if tok = token.TYPE.
// Parameter tok should be token.TYPE or token.VAR.
func (s TypeBuilder) CompleteAsDecl(tok token.Token) *ast.GenDecl {
	return &ast.GenDecl{
		Specs: []ast.Spec{s.CompleteAsSpec()},
		Tok:   tok,
	}
}

// CompleteAll returns type declaration and all associated methods.
func (s TypeBuilder) CompleteAll() []ast.Decl {
	r := make([]ast.Decl, 0, len(s.methods)+1)

	r = append(r, s.CompleteAsDecl(token.TYPE))
	for _, decl := range s.CompleteMethods() {
		r = append(r, decl)
	}

	return r
}
