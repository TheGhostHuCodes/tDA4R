lexer grammar XMLLexer;

// Default "mode": Everything outside of a tag.
OPEN: '<' -> pushMode(INSIDE);
COMMENT: '<!--' .*? '-->';
EntityRef: '&' [a-z]+ ';';
TEXT:
	~('<' | '&')+; // Maches any 16-bit char that is not < or &.

// Everything INSIDE of a tag.
mode INSIDE;

CLOSE: '>' -> popMode;
SLASH_CLOSE: '/>' -> popMode;
EQUALS: '=';
STRING: '"' .*? '"';
SlashName: '/' Name;
Name: ALPHA (ALPHA | DIGIT)*;
S: [ \t\r\n] -> skip;

fragment ALPHA: [a-zA-Z];

fragment DIGIT: [0-9];