package builders

import (
	"go/ast"
)

type StructBuilder struct {
	fields []*ast.Field
}

// EmptyStruct creates new empty struct{} declaration.
func EmptyStruct() *ast.StructType {
	return NewStructBuilder().Complete()
}

// NewStructBuilder creates new StructBuilder.
func NewStructBuilder() StructBuilder {
	return StructBuilder{}
}

// AddFields adds fields to struct definition.
func (s StructBuilder) AddFields(fields ...*ast.Field) StructBuilder {
	s.fields = append(s.fields, fields...)
	return s
}

// Complete returns struct definition.
// i.e struct{/* fields */}.
func (s StructBuilder) Complete() *ast.StructType {
	return &ast.StructType{
		Fields: &ast.FieldList{
			List: s.fields,
		},
	}
}

// CompleteAsLit returns a composite literal with built struct.
// i.e. struct{/* fields */}{/* elements */}
func (s StructBuilder) CompleteAsLit(elements ...ast.Expr) *ast.CompositeLit {
	return &ast.CompositeLit{
		Type: s.Complete(),
		Elts: elements,
	}
}
