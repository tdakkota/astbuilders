package builders_test

import (
	"go/printer"
	"go/token"
	"os"

	builders "github.com/tdakkota/astbuilders"
)

func ExampleImports() {
	imports := builders.Imports(
		builders.Import("fmt"),
		builders.NamedImport("builders", "github.com/tdakkota/astbuilders"),
	)

	printer.Fprint(os.Stdout, token.NewFileSet(), imports)
	// Output: import (
	// 	"fmt"
	// 	builders "github.com/tdakkota/astbuilders"
	// )
}
