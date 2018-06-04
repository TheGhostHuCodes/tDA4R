.PHONY: javabuild clean

csv-grammar := CSV
json-grammar := JSON
build-dir := build

javabuild: export CLASSPATH=".:/usr/local/Cellar/antlr/4.7.1/antlr-4.7.1-complete.jar:$$CLASSPATH"
javabuild:
	antlr4 -o $(build-dir) $(csv-grammar).g4
	antlr4 -o $(build-dir) $(json-grammar).g4
	javac $(build-dir)/*.java

clean:
	rm -rf $(build-dir)
