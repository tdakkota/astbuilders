package builders

import "go/ast"

type FunctionBuilder struct {
	name     *ast.Ident
	funcType *ast.FuncType
	recv     *ast.Field
	stmts    []ast.Stmt
	docs     *ast.CommentGroup
}

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

func (builder FunctionBuilder) AddParameters(params ...*ast.Field) FunctionBuilder {
	builder.funcType.Params.List = append(builder.funcType.Params.List, params...)
	return builder
}

func (builder FunctionBuilder) AddResults(results ...*ast.Field) FunctionBuilder {
	builder.funcType.Results.List = append(builder.funcType.Results.List, results...)
	return builder
}

func (builder FunctionBuilder) AddStmts(stmts ...ast.Stmt) FunctionBuilder {
	builder.stmts = append(builder.stmts, stmts...)
	return builder
}

func (builder FunctionBuilder) Body(f BodyFunc) FunctionBuilder {
	return builder.AddStmts(f(StatementBuilder{}).Complete()...)
}

func (builder FunctionBuilder) Recv(recv *ast.Field) FunctionBuilder {
	builder.recv = recv
	return builder
}

func (builder FunctionBuilder) CompleteAsLiteral() *ast.FuncLit {
	return &ast.FuncLit{
		Type: builder.funcType,
		Body: &ast.BlockStmt{
			List: builder.stmts,
		},
	}
}

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

func (builder FunctionBuilder) CompleteAsCall(args ...ast.Expr) *ast.CallExpr {
	lit := builder.CompleteAsLiteral()
	return &ast.CallExpr{
		Fun:  lit,
		Args: args,
	}
}
