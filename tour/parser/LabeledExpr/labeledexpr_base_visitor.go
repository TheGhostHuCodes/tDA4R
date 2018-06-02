// Code generated from LabeledExpr.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // LabeledExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLabeledExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLabeledExprVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitBlank(ctx *BlankContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}
