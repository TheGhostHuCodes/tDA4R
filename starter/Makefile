.PHONY: build clean

grammar := ArrayInit
build-dir := build
parser-dir := parser

gobuild:
	antlr4 -Dlanguage=Go -o $(parser-dir) $(grammar).g4
	go build -o $(build-dir)/Test test/Test.go
	go build -o $(build-dir)/Translate translate/Translate.go

javabuild: export CLASSPATH=".:/usr/local/Cellar/antlr/4.7.1/antlr-4.7.1-complete.jar:$$CLASSPATH"
javabuild:
	antlr4 -o $(build-dir) $(grammar).g4
	javac $(build-dir)/*.java

clean:
	rm -rf $(build-dir)
	rm -rf $(parser-dir)
