# astbuilders

[![Go](https://github.com/tdakkota/astbuilders/workflows/Go/badge.svg)](https://github.com/tdakkota/astbuilders/actions)
[![Documentation](https://godoc.org/github.com/tdakkota/astbuilders?status.svg)](https://pkg.go.dev/github.com/tdakkota/astbuilders)
[![codecov](https://codecov.io/gh/tdakkota/astbuilders/branch/master/graph/badge.svg)](https://codecov.io/gh/tdakkota/astbuilders)
[![license](https://img.shields.io/github/license/tdakkota/astbuilders.svg)](https://github.com/tdakkota/astbuilders/blob/master/LICENSE)

Go AST utility package 

## Install
```
go get github.com/tdakkota/astbuilders
```

## Examples

### Creating a function
```go
package main

import (
	"go/printer"
	"go/token"
	"os"

	"github.com/tdakkota/astbuilders"
)

func main() {
	node := builders.NewFunctionBuilder("main").
		Body(func(s builders.StatementBuilder) builders.StatementBuilder {
			return s.Expr(builders.CallPackage("fmt", "Println", builders.StringLit("Hello, world!")))
		}).
		CompleteAsDecl()

	printer.Fprint(os.Stdout, token.NewFileSet(), node)
}
```
prints
```go
func main() {
    fmt.Println("Hello, world!")
}
```

### `if err != nil`
```go
package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	
	"github.com/tdakkota/astbuilders"
)

func main() {
	errIdent := ast.NewIdent("err")
	nilIdent := ast.NewIdent("nil")
	cond := builders.NotEq(errIdent, nilIdent)

	s := builders.NewStatementBuilder()
	s = s.If(nil, cond, func(body builders.StatementBuilder) builders.StatementBuilder {
		return body.Return(errIdent)
	})

	stmts := s.Complete()
	node := stmts[0]
	printer.Fprint(os.Stdout, token.NewFileSet(), node)
}
```
prints
```go
if err != nil {
	return err
}
```

### Reverse from SliceTricks
```go
package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	
	"github.com/tdakkota/astbuilders"
)

func main() {
	s := builders.NewStatementBuilder()

	a := ast.NewIdent("a")
	left, right := ast.NewIdent("left"), ast.NewIdent("right")
	one := builders.IntegerLit(1)

	// left, right := 0, len(a)-1
	init := builders.Define(left, right)(builders.IntegerLit(0), builders.Sub(builders.Len(a), one))
	// left < right
	cond := builders.Less(left, right)
	// left, right = left+1, right-1
	post := builders.Assign(left, right)(token.ASSIGN)(builders.Add(left, one), builders.Sub(right, one))

	// a[left]
	indexLeft := builders.Index(a, left)
	// a[right]
	indexRight := builders.Index(a, right)

	// for $init; $cond; $post {
	// for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
	s = s.For(init, cond, post, func(loop builders.StatementBuilder) builders.StatementBuilder {
		// a[left], a[right] = a[right], a[left]
		loop = loop.AddStmts(builders.Swap(indexLeft, indexRight))
		return loop
	})

	stmts := s.Complete()
	node := stmts[0]
	printer.Fprint(os.Stdout, token.NewFileSet(), node)
}
```

prints

```go
for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
    a[left], a[right] = a[right], a[left]
}
```