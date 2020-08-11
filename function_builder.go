package builders

import "go/ast"

// FunctionBuilder is function definition builder.
type FunctionBuilder struct {
	name     *ast.Ident
	funcType *ast.FuncType
	recv     *ast.Field
	stmts    []ast.Stmt
	docs     *ast.CommentGroup
}

// NewFunctionBuilder creates new FunctionBuilder.
func NewFunctionBuilder(name string) FunctionBuilder {
	return FunctionBuilder{
		name: ast.NewIdent(name),
		funcType: &ast.FuncType{
			Params: &ast.FieldList{
				List: make([]*ast.Field, 0),
			},
			Results: &ast.FieldList{
				List: make([]*ast.Field, 0),
			},
		},
	}
}

// AddParameters adds elements to function parameters.
func (builder FunctionBuilder) AddParameters(params ...*ast.Field) FunctionBuilder {
	builder.funcType.Params.List = append(builder.funcType.Params.List, params...)
	return builder
}

// AddParameters adds elements to result tuple.
func (builder FunctionBuilder) AddResults(results ...*ast.Field) FunctionBuilder {
	builder.funcType.Results.List = append(builder.funcType.Results.List, results...)
	return builder
}

// AddStmts adds statements to function body.
func (builder FunctionBuilder) AddStmts(stmts ...ast.Stmt) FunctionBuilder {
	builder.stmts = append(builder.stmts, stmts...)
	return builder
}

// Body adds statements returned from BodyFunc to function body.
func (builder FunctionBuilder) Body(f BodyFunc) FunctionBuilder {
	return builder.AddStmts(f(StatementBuilder{}).Complete()...)
}

// Recv sets receiver of function.
func (builder FunctionBuilder) Recv(recv *ast.Field) FunctionBuilder {
	builder.recv = recv
	return builder
}

// CompleteAsLiteral returns built function as function literal.
func (builder FunctionBuilder) CompleteAsLiteral() *ast.FuncLit {
	return &ast.FuncLit{
		Type: builder.funcType,
		Body: &ast.BlockStmt{
			List: builder.stmts,
		},
	}
}

// CompleteAsDecl returns built function as function declaration.
func (builder FunctionBuilder) CompleteAsDecl() *ast.FuncDecl {
	var recv *ast.FieldList
	if builder.recv != nil {
		recv = &ast.FieldList{List: []*ast.Field{builder.recv}}
	}

	return &ast.FuncDecl{
		Doc:  builder.docs,
		Recv: recv,
		Name: builder.name,
		Type: builder.funcType,
		Body: &ast.BlockStmt{
			List: builder.stmts,
		},
	}
}

// CompleteAsCall returns built function as function literal call.
func (builder FunctionBuilder) CompleteAsCall(args ...ast.Expr) *ast.CallExpr {
	lit := builder.CompleteAsLiteral()
	return &ast.CallExpr{
		Fun:  lit,
		Args: args,
	}
}
