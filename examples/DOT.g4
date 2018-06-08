grammar DOT;

graph: STRICT? (GRAPH | DIGRAPH) id? '{' statement_list '}';
statement_list: (statement ';'?)*;
statement:
	node_statement
	| edge_statement
	| attribute_statement
	| id '=' id
	| subgraph;
attribute_statement: (GRAPH | NODE | EDGE) attribute_list;
attribute_list: ('[' a_list? ']')+;
a_list: (id ('=' id)? ','?)+;
edge_statement: (node_id | subgraph) edgeRHS attribute_list?;
edgeRHS: (edgeop (node_id | subgraph))+;
edgeop: '->' | '--';
node_statement: node_id attribute_list?;
node_id: id port?;
port: ':' id (':' id)?;
subgraph: (SUBGRAPH id?)? '{' statement_list '}';
id: ID | STRING | HTML_STRING | NUMBER;

STRICT: [Ss][Tt][Rr][Ii][Cc][Tt];
GRAPH: [Gg][Rr][Aa][Pp][Hh];
DIGRAPH: [Dd][Ii][Gg][Rr][Aa][Pp][Hh];
NODE: [Nn][Oo][Dd][Ee];
EDGE: [Ee][Dd][Gg][Ee];
SUBGRAPH: [Ss][Uu][Bb][Gg][Rr][Aa][Pp][Hh];

ID: LETTER (LETTER | DIGIT)*;
fragment LETTER: [a-zA-Z\u0080-\u00FF_];

NUMBER: '-' ('.' DIGIT+ | DIGIT+ ('.' DIGIT*)?);
fragment DIGIT: [0-9];

STRING: '"' ('\\"' | .)*? '"';

/**
 * "HTML strings, angle brackets must occur in matched pairs, and unescaped newlines are allowed.
 */

HTML_STRING: '<' (TAG | ~[<>])* '>';
fragment TAG: '<' .*? '>';

PREPROC: '#' .*? '\n' -> skip;
LINE_COMMENT: '//' .*? '\r'? '\n' -> skip;
WS: [ \t\n\r]+ -> skip;