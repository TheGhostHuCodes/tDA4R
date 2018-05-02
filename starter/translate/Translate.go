package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/TheGhostHuCodes/tDA4R/starter/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type arrayInitListener struct {
	*parser.BaseArrayInitListener
}

func (l *arrayInitListener) EnterInit(c *parser.InitContext) {
	fmt.Printf("\"")
}

func (l *arrayInitListener) ExitInit(c *parser.InitContext) {
	fmt.Printf("\"")
}

// EnterValue translates integers to 4-digit hexadecimal strings, prefixed with
// '\u'.
func (l *arrayInitListener) EnterValue(c *parser.ValueContext) {
	// Assumes no nested array initializers.
	i, _ := strconv.Atoi(c.INT().GetText())
	fmt.Printf("\\u%04x", i)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	is := antlr.NewInputStream(text)
	lexer := parser.NewArrayInitLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewArrayInitParser(tokens)
	tree := p.Init()

	var listener arrayInitListener
	walker := antlr.NewParseTreeWalker()
	walker.Walk(&listener, tree)
}
