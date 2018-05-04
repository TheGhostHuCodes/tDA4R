package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TheGhostHuCodes/tDA4R/tour/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func readFromStdin() (*antlr.InputStream, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	return antlr.NewInputStream(string(bytes)), nil
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

	lexer := parser.NewExprLexer(s)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewExprParser(tokens)
	tree := p.Prog()
	fmt.Printf(tree.ToStringTree([]string{}, p))
}
