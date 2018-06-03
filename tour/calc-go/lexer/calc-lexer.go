package main

import (
	"fmt"

	"github.com/TheGhostHuCodes/tDA4R/tour/parser/Calc"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	is := antlr.NewInputStream("1 + 2 * 3")
	lexer := parser.NewCalcLexer(is)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
