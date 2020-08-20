package builders

import "go/ast"

type InterfaceBuilder struct {
	fields []*ast.Field
}

// EmptyInterface creates new empty interface{} declaration.
func EmptyInterface() *ast.InterfaceType {
	return NewInterfaceBuilder().Complete()
}

// NewInterfaceBuilder creates new InterfaceBuilder.
func NewInterfaceBuilder() InterfaceBuilder {
	return InterfaceBuilder{}
}

// AddMethod adds method to interface's method list.
func (s InterfaceBuilder) AddMethod(name string, typ *ast.FuncType) InterfaceBuilder {
	s.fields = append(s.fields, Param(ast.NewIdent(name))(typ))
	return s
}

// BuildMethod builds function type using callback and adds method to interface's method list.
func (s InterfaceBuilder) BuildMethod(name string, cb func(FunctionBuilder) FunctionBuilder) InterfaceBuilder {
	f := NewFunctionBuilder(name)
	return s.AddMethod(name, cb(f).CompleteAsType())
}

// Complete returns interface definition.
// i.e interface{/* methods */}.
func (s InterfaceBuilder) Complete() *ast.InterfaceType {
	return &ast.InterfaceType{
		Methods: &ast.FieldList{List: s.fields},
	}
}
