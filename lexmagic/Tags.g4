grammar Tags;

file: (TAG | ENTITY | TEXT | CDATA)*;

COMMENT: '<!--' .*? '-->' -> skip;
CDATA: '<![CDATA[' .*? ']]>';
TAG: '<' .*? '>'; // Must come after other tag-like structures
ENTITY: '&' .*? ';';
TEXT:
	~[<&]+; // Any sequence of characters except < and & characters.