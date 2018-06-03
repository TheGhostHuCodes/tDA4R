package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TheGhostHuCodes/tDA4R/tour/parser/Java"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func readFromStdin() (*antlr.InputStream, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	return antlr.NewInputStream(string(bytes)), nil
}

type javaListener struct {
	*parser.BaseJavaListener
	p *parser.JavaParser
}

func newExtractInterfaceListener(p *parser.JavaParser) *javaListener {
	return &javaListener{p: p}
}
func (l *javaListener) EnterClassDeclaration(c *parser.ClassDeclarationContext) {
	fmt.Printf("interface I%s {\n", c.Identifier())
}
func (l *javaListener) ExitClassDeclaration(c *parser.ClassDeclarationContext) {
	fmt.Printf("}\n")
}
func (l *javaListener) EnterMethodDeclaration(c *parser.MethodDeclarationContext) {
	tokens := l.p.GetTokenStream()
	dataType := "void"
	if c.Datatype() != nil {
		dataType = tokens.GetTextFromRuleContext(c.Datatype())
	}

	args := tokens.GetTextFromInterval(c.FormalParameters().GetSourceInterval())
	fmt.Printf("\t%s %s%s;\n", dataType, c.Identifier(), args)
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

	lexer := parser.NewJavaLexer(s)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaParser(tokens)

	tree := p.CompilationUnit()

	fmt.Printf(tree.ToStringTree([]string{}, p))
	fmt.Printf("\n\n")
	antlr.ParseTreeWalkerDefault.Walk(&javaListener{p: p}, tree)
}
