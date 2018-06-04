grammar CSV;

file: header row+;
header: row;
row: field (',' field)* '\r'? '\n';
field: TEXT | STRING |;

TEXT: ~[,\n\r"]+;
STRING: '"' ('""' | ~'"')* '"';