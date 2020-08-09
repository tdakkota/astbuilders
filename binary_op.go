package builders

import (
	"go/ast"
	"go/token"
)

func BinaryOp(x ast.Expr, tok token.Token, y ast.Expr) *ast.BinaryExpr {
	return &ast.BinaryExpr{
		X:  x,
		Op: tok,
		Y:  y,
	}
}

// Arithmetic

func Add(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.ADD, y)
}

func Sub(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.SUB, y)
}

func Mul(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.MUL, y)
}

func Div(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.QUO, y)
}

func Rem(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.REM, y)
}

// Compare

func Eq(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.EQL, y)
}

func NotEq(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.NEQ, y)
}

func Greater(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.GTR, y)
}

func GreaterOrEq(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.GEQ, y)
}

func Less(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.LSS, y)
}

func LessOrEq(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.LEQ, y)
}

// Login

func And(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.LAND, y)
}

func Or(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.LOR, y)
}

// Binary

func BAnd(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.AND, y)
}

func BOr(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.OR, y)
}

func Xor(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.XOR, y)
}

func AddNot(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.AND_NOT, y)
}

func ShiftLeft(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.SHL, y)
}

func ShiftRight(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.SHR, y)
}

// Channel

// Send returns x <- y expression.
func Send(x ast.Expr, y ast.Expr) *ast.BinaryExpr {
	return BinaryOp(x, token.ARROW, y)
}
