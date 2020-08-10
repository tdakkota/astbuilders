package builders_test

import (
	"go/printer"
	"go/token"
	"os"

	"github.com/tdakkota/astbuilders"
)

func Example_main() {
	node := builders.NewFunctionBuilder("main").
		Body(func(s builders.StatementBuilder) builders.StatementBuilder {
			return s.Expr(builders.CallPackage("fmt", "Println", builders.StringLit("Hello, world!")))
		}).
		CompleteAsDecl()

	printer.Fprint(os.Stdout, token.NewFileSet(), node)
	// Output: func main() {
	// 	fmt.Println("Hello, world!")
	// }
}
