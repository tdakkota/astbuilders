package builders

import (
	"go/ast"
	"go/token"
)

// BinaryOp returns binary expression with given token.
func BinaryOp(x ast.Expr, tok token.Token, y ast.Expr) *ast.BinaryExpr {
	return &ast.BinaryExpr{
		X:  x,
		Op: tok,
		Y:  y,
	}
}

// Arithmetic

// Add returns x + y expression.
func Add(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.ADD, y)
}

// Sub returns x - y expression.
func Sub(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.SUB, y)
}

// Mul returns x * y expression.
func Mul(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.MUL, y)
}

// Div returns x / y expression.
func Div(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.QUO, y)
}

// Rem returns x % y expression.
func Rem(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.REM, y)
}

// Compare

// Eq returns x == y expression.
func Eq(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.EQL, y)
}

// NotEq returns x != y expression.
func NotEq(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.NEQ, y)
}

// Greater returns x > y expression.
func Greater(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.GTR, y)
}

// GreaterOrEq returns x >= y expression.
func GreaterOrEq(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.GEQ, y)
}

// Less returns x < y expression.
func Less(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.LSS, y)
}

// LessOrEq returns x <= y expression.
func LessOrEq(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.LEQ, y)
}

// Logic

// And returns x && y expression.
func And(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.LAND, y)
}

// And returns x || y expression.
func Or(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.LOR, y)
}

// Bitwise

// BAnd returns x & y expression.
func BAnd(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.AND, y)
}

// BOr returns x | y expression.
func BOr(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.OR, y)
}

// BOr returns x ^ y expression.
func Xor(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.XOR, y)
}

// AddNot returns x &^ y expression.
func AddNot(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.AND_NOT, y)
}

// ShiftLeft returns x << y expression.
func ShiftLeft(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.SHL, y)
}

// ShiftRight returns x >> y expression.
func ShiftRight(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.SHR, y)
}

// Channel

// Send returns x <- y expression.
func Send(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.ARROW, y)
}
