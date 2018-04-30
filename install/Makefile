.PHONY: build clean

build: export CLASSPATH=".:/usr/local/Cellar/antlr/4.7.1/antlr-4.7.1-complete.jar:$$CLASSPATH"
build:
	antlr4 -o build Hello.g4
	javac build/*.java

clean:
	rm -rf build
