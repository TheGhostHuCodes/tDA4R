.PHONY: gobuild javabuild clean

grammar1 := Expr
build-dir := build
parser-dir := parser

gobuild:
	antlr4 -Dlanguage=Go -o $(parser-dir)/$(grammar1) $(grammar1).g4
	go build -o $(build-dir)/ExprJoyRide ExprJoyRide.go

javabuild: export CLASSPATH=".:/usr/local/Cellar/antlr/4.7.1/antlr-4.7.1-complete.jar:$$CLASSPATH"
javabuild:
	antlr4 -o $(build-dir) $(grammar1).g4
	javac $(build-dir)/*.java

clean:
	rm -rf $(build-dir)
	rm -rf $(parser-dir)
