package main

import (
	"fmt"

	"github.com/TheGhostHuCodes/tDA4R/tour/parser/Calc"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcListener struct {
	*parser.BaseCalcListener
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	fmt.Printf("Hit ExitMulDiv\n")
}
func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	fmt.Printf("Hit ExitAddSub\n")
}
func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	fmt.Printf("Hit ExitNumber\n")
}

func main() {
	is := antlr.NewInputStream("1 + 2 * 3")
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewCalcParser(stream)
	tree := p.Start()
	antlr.ParseTreeWalkerDefault.Walk(&calcListener{}, tree)
}
