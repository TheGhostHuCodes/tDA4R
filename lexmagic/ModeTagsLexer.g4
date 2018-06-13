lexer grammar ModeTagsLexer;

// Default mode rules (the SEA)
OPEN: '<' -> mode(ISLAND); // Switch to ISLAND mode.
TEXT: ~'<'+; // Clump all text together.

mode ISLAND;
CLOSE: '>' -> mode(DEFAULT_MODE); // bak to SEA mode.
SLASH: '/';
ID: [a-zA-Z]+; // match/send ID in tag to parser.