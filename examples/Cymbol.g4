grammar Cymbol;

file: (functionDeclaration | variableDeclaration)+;

variableDeclaration: type ID ('=' expression)? ';';
type: 'float' | 'int' | 'void'; // user-defined types

functionDeclaration:
	type ID '(' formalParameters? ')' block; // "void f(int x) {...}"
formalParameters: formalParameter (',' formalParameter)*;
formalParameter: type ID;

block: '{' statement* '}'; // possibly empty statement block
statement:
	block
	| variableDeclaration
	| 'if' expression 'then' statement ('else' statement)?
	| 'return' expression? ';'
	| expression '=' expression ';' // assignment
	| expression ';'; // function call

expression:
	ID '(' expressionList? ')' // func call like f(), f(x), f(1, 2)
	| ID '[' expression ']' // array index like a[i], a[i][j]
	| '-' expression // unary minus
	| '!' expression // boolean not
	| expression '*' expression
	| expression ('+' | '-') expression
	| expression '==' expression
	| ID
	| INT
	| '(' expression ')';
expressionList: expression (',' expression)*; // argument list

ID: LETTER (LETTER | [0-9])*;
fragment LETTER: [a-zA-Z];

INT: [0-9]+;
WS: [ \t\n\r]+ -> skip;
LINE_COMMENT: '//' .*? '\r'? '\n' -> skip;