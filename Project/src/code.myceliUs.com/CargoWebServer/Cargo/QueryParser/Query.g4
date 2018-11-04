grammar Query;

AND     : ',' | 'and' | '&''&' ;
OR      : 'or' | '|''|' ;
INT     : '-'? [0-9]+ ;
DOUBLE  : '-'? [0-9]+'.'[0-9]+ ;
ID      : [a-zA-Z_][a-zA-Z_0-9]* ;
STRING  :  '"' (~["])* '"' | '\'' (~['])* '\'';
BOOLEAN : 'T''R''U''E'|'F''A''L''S''E'|'t''r''u''e'|'f''a''l''s''e';
EQ      : '=' '=' ? ;
LE      : '<=' ;
GE      : '>=' ;
NE      : '!=' ;
LT      : '<' ;
GT      : '>' ;
MAYBE   : '~=';
SEP     : '.' ;
WS      : [ \t\r\n]+ -> skip ;

query 
      : expression
      ;

expression
    : expression AND expression # AndExpression
    | expression OR expression  # OrExpression
    | predicate                 # PredicateExpression
    | '(' expression ')'        # BracketExpression
    ;

reference : element (SEP element)* ;

element : ID ;

predicate 
    :reference operator value
    ;

operator
    : LE
    | GE
    | NE
    | LT
    | GT
    | EQ
    | MAYBE
    ;

value
   : INT          # IntegerValue
   | DOUBLE       # DoubleValue
   | BOOLEAN      # BooleanValue
   | STRING       # StringValue
   | ID           # StringValue
   ;

