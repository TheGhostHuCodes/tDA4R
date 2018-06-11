grammar R;

prog: (expression_or_assignment (';' | NL) | NL)* EOF;

expression_or_assignment:
	expression ('<-' | '=' | '<<-') expression_or_assignment
	| expression;

expression:
	expression '[[' sublist ']' ']' // '[[' follows R's yacc grammar
	| expression '[' sublist ']'
	| expression ('::' | ':::') expression
	| expression ('$' | '@') expression
	| <assoc = right> expression '^' expression
	| ('-' | '+') expression
	| expression ':' expression
	| expression USER_OP expression // anything wrapped in %: '%' .* '%'
	| expression ('*' | '/') expression
	| expression ('+' | '-') expression
	| expression ('>' | '>=' | '<' | '<=' | '==' | '!=') expression
	| '!' expression
	| expression ('&' | '&&') expression
	| expression ('/' | '//') expression
	| '~' expression
	| expression '~' expression
	| expression ('->' | '->>' | ':=') expression
	| 'function' '(' formlist? ')' expression // define function
	| expression '(' sublist ')' // call function
	| '{' expression_list '}' // compound statement
	| 'if' '(' expression ')' expression
	| 'if' '(' expression ')' expression 'else' expression
	| 'for' '(' ID 'in' expression ')' expression
	| 'while' '(' expression ')' expression
	| 'repeat' expression
	| '?' expression // get help on expression, usually string or ID
	| 'next'
	| 'break'
	| '(' expression ')'
	| ID
	| STRING
	| HEX
	| INT
	| FLOAT
	| COMPLEX
	| 'NULL'
	| 'NA'
	| 'Inf'
	| 'NaN'
	| 'TRUE'
	| 'FALSE';

expression_list:
	expression_or_assignment (
		(';' | NL) expression_or_assignment?
	)*
	|;

formlist: form (',' form)*;
form: ID | ID '=' expression | '...';

sublist: sub (',' sub)*;
sub:
	expression
	| ID '='
	| ID '=' expression
	| STRING '='
	| STRING '=' expression
	| 'NULL' '='
	| 'NULL' '=' expression
	| '...'
	|;

ID:
	'.' (LETTER | '_' | '.') (LETTER | DIGIT | '_' | '.')*
	| LETTER (LETTER | DIGIT | '_' | '.')*;
fragment LETTER: [a-zA-Z];
fragment DIGIT: [0-9];

HEX: '0' ('x' | 'X') HEXDIGIT+ [Ll]?;

INT: DIGIT+ [Ll]?;

FLOAT:
	DIGIT+ '.' DIGIT* EXP? [Ll]?
	| DIGIT+ EXP? [Ll]?
	| '.' DIGIT+ EXP? [Ll]?;

EXP: ('E' | 'e') ('+' | '-')? INT;

COMPLEX: INT 'i' | FLOAT 'i';

STRING: '"' ( ESC | ~[\\"])*? '"';
fragment ESC:
	'\\' ([abtnfrv] | '"' | '\'')
	| UNICODE_ESCAPE
	| HEX_ESCAPE
	| OCTAL_ESCAPE;

fragment UNICODE_ESCAPE:
	'\\' 'u' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
	| '\\' 'u' '{' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT '}';
fragment HEX_ESCAPE: '\\' HEXDIGIT HEXDIGIT?;

fragment HEXDIGIT: ('0' ..'9' | 'a' ..'f' | 'A' ..'F');

fragment OCTAL_ESCAPE:
	'\\' [0-3] [0-7] [0-7]
	| '\\' [0-7] [0-7]
	| '\\' [0-7];

USER_OP: '%' .*? '%';

// Match both UNIX and Windows newlines.
NL: '\r'? '\n';

WS: [ \t]+ -> skip;