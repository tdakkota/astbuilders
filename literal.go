package builders

import (
	"go/ast"
	"go/token"
	"strconv"
)

// StringLit creates string literal.
func StringLit(value string) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: strconv.Quote(value),
	}
}

// CharLit creates rune/character literal.
func CharLit(value rune) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.CHAR,
		Value: strconv.QuoteRune(value),
	}
}

// IntegerLit creates integer literal.
func IntegerLit(value int) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.INT,
		Value: strconv.Itoa(value),
	}
}

// FloatLit creates floating point literal.
func FloatLit(value float64) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.FLOAT,
		Value: strconv.FormatFloat(value, 'f', -1, 64),
	}
}
