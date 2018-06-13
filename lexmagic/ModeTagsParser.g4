parser grammar ModeTagsParser;

options {
	// Use tokens from ModeTagsLexer.g4
	tokenVocab = ModeTagsLexer;
}

file: (tag | TEXT)*;

tag: '<' ID '>' | '<' '/' ID '>';