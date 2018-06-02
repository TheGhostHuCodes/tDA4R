package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/TheGhostHuCodes/tDA4R/tour/parser/LabeledExpr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func readFromStdin() (*antlr.InputStream, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	return antlr.NewInputStream(string(bytes)), nil
}

type evalVisitor struct {
	*parser.BaseLabeledExprVisitor
	memory map[string]int
}

func newEvalVisitor() *evalVisitor {
	return &evalVisitor{memory: make(map[string]int)}
}
func (l *evalVisitor) VisitAssign(c *parser.AssignContext) int {
	id := c.ID().GetText()
	value := l.Visit(c.Expr()).(int)
	l.memory[id] = value
	return value
}
func (l *evalVisitor) VisitPrintExpr(c *parser.PrintExprContext) {
	value := l.Visit(c.Expr()).(int)
	fmt.Println(value)
}
func (l *evalVisitor) VisitInt(c *parser.IntContext) int {
	i, _ := strconv.Atoi(c.INT().GetText())
	return i
}
func (l *evalVisitor) VisitId(c *parser.IdContext) int {
	id := c.ID().GetText()
	value, present := l.memory[id]
	if present {
		return value
	}
	return 0
}
func (l *evalVisitor) VisitMulDiv(c *parser.MulDivContext) int {
	left := l.Visit(c.Expr(0)).(int)
	right := l.Visit(c.Expr(1)).(int)
	if c.GetOp().GetTokenType() == parser.LabeledExprParserMUL {
		return left * right
	}
	return left / right
}
func (l *evalVisitor) VisitAddSub(c *parser.AddSubContext) int {
	left := l.Visit(c.Expr(0)).(int)
	right := l.Visit(c.Expr(1)).(int)
	if c.GetOp().GetTokenType() == parser.LabeledExprParserADD {
		return left + right
	}
	return left - right
}
func (l *evalVisitor) VisitParens(c *parser.ParensContext) interface{} {
	return l.Visit(c.Expr())
}

func main() {
	var s antlr.CharStream
	var err error
	if len(os.Args) == 1 {
		s, err = readFromStdin()
		if err != nil {
			panic(fmt.Sprintf("ERROR: %s\n", err.Error()))
		}
	} else if len(os.Args) == 2 {
		s, err = antlr.NewFileStream(os.Args[1])
		if err != nil {
			panic(fmt.Sprintf("ERROR: %s\n", err.Error()))
		}
	} else {
		panic(fmt.Sprintf("Expected zero or one commandline argument(s)"))
	}

	lexer := parser.NewLabeledExprLexer(s)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewLabeledExprParser(tokens)
	tree := p.Prog()
	// fmt.Printf(tree.ToStringTree([]string{}, p))
	eval := newEvalVisitor()
	//	fmt.Printf("%s\n", eval.memory)
	eval.Visit(tree)
}
