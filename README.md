# The Definitive ANTLR 4 Reference, 2nd Edition

My workspace containing code listings from the book ["The Definitive ANTLR 4
Reference, 2nd
Edition"](https://pragprog.com/book/tpantlr2/the-definitive-antlr-4-reference),
by Terence Parr.

The code examples in the book are written in Java. However, since the release
of ANTLR 4, Go is an officially supported target language. I wanted to write in
Go, so examples have been converted to Go.

## Setup

Install ANTLR using homebrew:

```shell
brew install antlr
```

When using homebrew to install ANTLR, you will need to export a new class path
to be able to use antlr to compile java code:

```shell
export CLASSPATH=".:/usr/local/Cellar/antlr/4.7.1/antlr-4.7.1-complete.jar:$CLASSPATH"
```
