// Code generated from Expr.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExprListener is a complete listener for a parse tree produced by ExprParser.
type ExprListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
