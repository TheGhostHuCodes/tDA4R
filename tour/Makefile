.PHONY: gobuild javabuild pybuild clean

grammar1 := Expr
grammar2 := LabeledExpr
build-dir := build
parser-dir := parser

gobuild:
	antlr4 -Dlanguage=Go -o $(parser-dir)/$(grammar1) $(grammar1).g4
	go build -o $(build-dir)/ExprJoyRide exprJoyRide/ExprJoyRide.go

pybuild:
	antlr4 -Dlanguage=Python3 -no-listener -visitor -o calc-py/$(grammar2) $(grammar2).g4

javabuild: export CLASSPATH=".:/usr/local/Cellar/antlr/4.7.1/antlr-4.7.1-complete.jar:$$CLASSPATH"
javabuild:
	antlr4 -o $(build-dir) $(grammar1).g4
	antlr4 -no-listener -visitor -o $(build-dir) $(grammar2).g4
	javac $(build-dir)/*.java

clean:
	rm -rf $(build-dir)
	rm -rf $(parser-dir)
