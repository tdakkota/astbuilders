package builders

import (
	"go/ast"
	"go/token"
	"strconv"
)

func StringLit(value string) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: strconv.Quote(value),
	}
}

func CharLit(value rune) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.CHAR,
		Value: strconv.QuoteRune(value),
	}
}

func IntegerLit(value int) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.INT,
		Value: strconv.Itoa(value),
	}
}

func FloatLit(value float64) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.INT,
		Value: strconv.FormatFloat(value, 'f', -1, 64),
	}
}
