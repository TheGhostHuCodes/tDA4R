.PHONY: javabuild clean

tags-grammar := Tags
build-dir := build

javabuild: export CLASSPATH=".:/usr/local/Cellar/antlr/4.7.1/antlr-4.7.1-complete.jar:$$CLASSPATH"
javabuild:
	antlr4 -o $(build-dir) $(tags-grammar).g4
	javac $(build-dir)/*.java

clean:
	rm -rf $(build-dir)
