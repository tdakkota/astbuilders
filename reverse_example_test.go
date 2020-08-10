package builders_test

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	"github.com/tdakkota/astbuilders"
)

func Example_Reverse() {
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

	// Output: for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
	// 	a[left], a[right] = a[right], a[left]
	// }
}
