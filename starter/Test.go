package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TheGhostHuCodes/tDA4R/starter/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	is := antlr.NewInputStream(text)
	lexer := parser.NewArrayInitLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewArrayInitParser(tokens)
	tree := p.Init()
	fmt.Printf(tree.ToStringTree([]string{}, p))
}
