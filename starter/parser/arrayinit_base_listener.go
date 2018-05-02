// Code generated from ArrayInit.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ArrayInit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseArrayInitListener is a complete listener for a parse tree produced by ArrayInitParser.
type BaseArrayInitListener struct{}

var _ ArrayInitListener = &BaseArrayInitListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseArrayInitListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseArrayInitListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseArrayInitListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseArrayInitListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInit is called when production init is entered.
func (s *BaseArrayInitListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseArrayInitListener) ExitInit(ctx *InitContext) {}

// EnterValue is called when production value is entered.
func (s *BaseArrayInitListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseArrayInitListener) ExitValue(ctx *ValueContext) {}
