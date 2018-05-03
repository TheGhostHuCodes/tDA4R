package main

import (
	"fmt"
	"os"

	"github.com/TheGhostHuCodes/tDA4R/tour/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Sprintf("Expected one commandline argument"))
	}
	arg := os.Args[1]

	fs, err := antlr.NewFileStream(arg)
	if err != nil {
		panic(fmt.Sprintf("ERROR: %s\n", err.Error()))
	}

	lexer := parser.NewExprLexer(fs)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewExprParser(tokens)
	tree := p.Prog()
	fmt.Printf(tree.ToStringTree([]string{}, p))
}
