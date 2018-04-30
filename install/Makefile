.PHONY: build clean

grammar := Hello
build-dir := build

build: export CLASSPATH=".:/usr/local/Cellar/antlr/4.7.1/antlr-4.7.1-complete.jar:$$CLASSPATH"
build:
	antlr4 -o $(build-dir) $(grammar).g4
	javac $(build-dir)/*.java

clean:
	rm -rf $(build-dir)
