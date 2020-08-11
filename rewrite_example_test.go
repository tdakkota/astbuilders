package builders_test

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"

	"github.com/tdakkota/astbuilders"
)

const code = `
package main

import "fmt"

func main() {
	fmt.Print("Hello, ")
}
`

func Example_rewrite() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, decl := range file.Decls {
		if v, ok := decl.(*ast.FuncDecl); ok && v.Name.Name == "main" {
			s := builders.NewStatementBuilder()
			s = s.AddStmts(v.Body.List...)
			s = s.Expr(builders.CallPackage("fmt", "Println", builders.StringLit("world!")))
			v.Body = s.CompleteAsBlock()
		}
	}

	printer.Fprint(os.Stdout, fset, file)
	// Output: package main
	//
	// import "fmt"
	//
	// func main()	{ fmt.Print("Hello, "); fmt.Println("world!") }
}
