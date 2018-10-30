// Define the grammar for JSONiq
grammar jsoniq;

// Program is the highest-level entity in JSONiq and
// can either be a library module (a collection of functions) or a
// main module that includes a query expression that will be
// evaluated. The main module can also contain a number of function
// definitions.

program :( (K_JSONIQ ((K_ENCODING StringLiteral) |
            (K_VERSION StringLiteral (K_ENCODING StringLiteral)?)) ';') ?
          (libraryModule | mainModule)) ;

// The Main Module consists of the Prolog and the query expression

mainModule : prolog expr ;

// The Library Module declares a namespace and the Prolog - which is a collection of
// settings and function and variable declarations

libraryModule : K_MODULE K_NAMESPACE NCName '=' uriLiteral ';' prolog ;

// The Prolog, it contains settings (like collation declaration) and a list of function and variable
// declarations 

prolog   : ((defaultCollationDecl | orderingModeDecl | emptyOrderDecl | decimalFormatDecl | moduleImport) ';')* 
             ((functionDecl | varDecl) ';')* ;

// Declaration of default collation

defaultCollationDecl : K_DECLARE K_DEFAULT K_COLLATION uriLiteral ;

// Ordering model 

orderingModeDecl : K_DECLARE K_ORDERING (K_ORDERED | K_UNORDERED) ;

// Empty ordering declaration

emptyOrderDecl : K_DECLARE K_DEFAULT K_ORDER K_EMPTY (K_GREATEST | K_LEAST) ;

// Declaration of decimal format

decimalFormatDecl : K_DECLARE ((K_DECIMAL_FORMAT (NCName ':')? NCName) | (K_DEFAULT K_DECIMAL_FORMAT)) (dfPropertyName '=' StringLiteral)* ;

// Decimal format property name

dfPropertyName : K_DECIMAL_SEPARATOR | K_GROUPING_SEPARATOR | K_INFINITY | K_MINUS_SIGN | K_NAN | K_PERCENT | K_PER_MILLE | K_ZERO_DIGIT | K_DIGIT | K_PATTERN_SEPARATOR ;

// Module import directive

moduleImport : K_IMPORT K_MODULE (K_NAMESPACE NCName '=')? uriLiteral (K_AT uriLiteral (',' uriLiteral)*)? ;

// Variable declaration

varDecl  : K_DECLARE K_VARIABLE varRef (K_AS sequenceType)? ((':=' exprSingle) | (K_EXTERNAL (':=' exprSingle)?)) ;

// Function declaration

functionDecl : K_DECLARE K_FUNCTION (NCName ':')? NCName '(' paramList? ')' (K_AS sequenceType)? ('{' expr '}' | K_EXTERNAL) ;

// Parameter list for a function declaration

paramList : '$' NCName (K_AS sequenceType)? (',' '$' NCName (K_AS sequenceType)?)* ;

// Expression sequence

expr     : exprSingle (',' exprSingle)* ;

// Single expression

exprSingle : flowrExpr
           | quantifiedExpr
           | switchExpr
           | typeswitchExpr
           | ifExpr
           | tryCatchExpr
           | orExpr ;

// Flowr clause expression

flowrExpr : (forClause | letClause)
             (forClause | letClause | whereClause | groupByClause | orderByClause | countClause)*
             K_RETURN exprSingle ;

// For clause

forClause : K_FOR varRef (K_AS sequenceType)? (K_ALLOWING K_EMPTY)? (K_AT varRef)? K_IN exprSingle (',' varRef (K_AS sequenceType)? (K_ALLOWING K_EMPTY)? (K_AT varRef)? K_IN exprSingle)* ;

// Let clause

letClause : K_LET varRef (K_AS sequenceType)? ':=' exprSingle (',' varRef (K_AS sequenceType)? ':=' exprSingle)* ;

// Count clause

countClause : K_COUNT varRef ;

// Where clause

whereClause : K_WHERE exprSingle ;

// Group-by clause

groupByClause : K_GROUP K_BY varRef ((K_AS sequenceType)? ':=' exprSingle)? (K_COLLATION uriLiteral)? (',' varRef ((K_AS sequenceType)? ':=' exprSingle)? (K_COLLATION uriLiteral)?)* ;

// Order-by clause

orderByClause : ((K_ORDER K_BY) | (K_STABLE K_ORDER K_BY)) exprSingle (K_ASCENDING | K_DESCENDING)? (K_EMPTY (K_GREATEST | K_LEAST))? (K_COLLATION uriLiteral)? (',' exprSingle (K_ASCENDING | K_DESCENDING )? (K_EMPTY (K_GREATEST | K_LEAST))? (K_COLLATION uriLiteral)?)* ;

// Universal or existential quantification

quantifiedExpr : (K_SOME | K_EVERY ) varRef (K_AS sequenceType)? K_IN exprSingle (',' varRef (K_AS sequenceType)? K_IN exprSingle)* K_SATISFIES exprSingle ;

// Switch expression 

switchExpr : K_SWITCH '(' expr ')' switchCaseClause+ K_DEFAULT K_RETURN exprSingle ;

// Case of the switch expression

switchCaseClause : (K_CASE exprSingle)+ K_RETURN exprSingle ;

// Type swtich inside the switch expression

typeswitchExpr : K_TYPESWITCH '(' expr ')' caseClause+ K_DEFAULT (varRef)? K_RETURN exprSingle ;

// Case clause inside the switch expression 

caseClause : K_CASE (varRef K_AS)? sequenceType ('|' sequenceType)* K_RETURN exprSingle ;

// If then else expression

ifExpr   : K_IF '(' expr ')' K_THEN exprSingle K_ELSE exprSingle ;

//Try-catch expression

tryCatchExpr : K_TRY '{' expr '}' K_CATCH '*' '{' expr '}' ;

// Logical and string expressions

orExpr   : andExpr ( K_OR andExpr )* ;
andExpr  : notExpr ( K_AND notExpr )* ;
notExpr  : K_NOT ? comparisonExpr2 ;
comparisonExpr2 : comparisonExpr ( ('=' | '!=' | '<' | '<=' | '>' | '=>') comparisonExpr )? ;
comparisonExpr : stringConcatExpr ( (K_EQ | K_NE | K_LT | K_LE | K_GT | K_GE) stringConcatExpr )? ;
stringConcatExpr : rangeExpr ( '||' rangeExpr )* ;

// Range expression

rangeExpr : additiveExpr ( K_TO additiveExpr )? ;

// Arithmetic expressions

additiveExpr : multiplicativeExpr ( ('+' | '-') multiplicativeExpr )* ;
multiplicativeExpr : instanceofExpr ( ('*' | K_DIV | K_IDIV | K_MOD ) instanceofExpr )* ;

// Type checking and manipulation expressions

instanceofExpr : treatExpr ( K_INSTANCE K_OF sequenceType )? ;
treatExpr : castableExpr ( K_TREAT K_AS sequenceType )? ;
castableExpr : castExpr ( K_CASTABLE K_AS atomicType '?'? )? ;
castExpr : unaryExpr ( K_CAST K_AS atomicType '?'? )? ;

// Unary arithmetic

unaryExpr : ('-' | '+')* simpleMapExpr ;
simpleMapExpr : postfixExpr ('!' postfixExpr)* ;

// Postfix expressions

postfixExpr : primaryExpr (predicate | objectLookup | arrayLookup | arrayUnboxing)* ;

// Predicate condition

predicate : '[' expr ']' ;

// Object lookup

objectLookup : '.' ( StringLiteral | NCName | parenthesizedExpr | varRef | contextItemExpr ) ;

// Array lookup and unboxing

arrayLookup : '[' '[' expr ']' ']' ;
arrayUnboxing : '[' ']' ;

// Atomic expressions

primaryExpr : literal
           | varRef
           | parenthesizedExpr
           | contextItemExpr
           | functionCall
           | orderedExpr
           | unorderedExpr
           | objectConstructor
           | arrayConstructor ;


// Variable reference, may include a namespace

varRef   : '$' (NCName ':')? NCName ;

parenthesizedExpr : '(' expr? ')' ;
contextItemExpr : '$$' ;

// Ordered and unordered modifiers

orderedExpr : K_ORDERED '{' expr '}' ;
unorderedExpr : K_UNORDERED '{' expr '}' ;

// Function call

functionCall : (NCName ':')? NCName argumentList ;

// Function call arguments:

argumentList: argument *;

// Function call argument

argument : exprSingle | '?' ;

// Object constructor

objectConstructor :  '{' ( pairConstructor (',' pairConstructor)* )? '}'
         | '{|' expr '|}' ;

// Pair constructoror

pairConstructor :  ( exprSingle | NCName ) (':' | '?') exprSingle ;

// Array constructor

arrayConstructor :  '[' expr? ']' ;


// Type declarations

sequenceType : '(' ')'
           | itemType ('?' | '*' | '+')? ;

itemType : K_ITEM
           | jsonItemTest
           | atomicType ;

jsonItemTest : K_OBJECT
           | K_ARRAY
           | K_JSON_ITEM ;

// Atomic types. The standard specifices 'string','boolean','integer','decimal','double', but also 
// mentions that all XQuery atomics are included

atomicType : K_ATOMIC | K_STRING | K_INTEGER
        | K_DECIMAL | K_DOUBLE | K_BOOLEAN
	| K_LONG | K_SHORT | K_BYTE
	| K_FLOAT | K_DATE | K_DATETIME | K_DATETIMESTAMP
	| K_GDAY | K_GMONTH | K_GMONTHDAY | K_GYEAR | K_GYEARMONTH
	| K_TIME | K_DURATION | K_DAYTIMEDURATION | K_YEARMONTHDURATION
	| K_BASE64NIBARY | K_HEXBINARY | K_ANYURI
        | K_NULL; 

// URI Literal

uriLiteral : StringLiteral ;

//*************************************
//  TOKENS
//*************************************

// Keywords. Since keywords in JSONiq are case-insensitive, we have
// to create lexical rules for them, given below (this is from ANTLR4 book,
// seems like there is no better way to do this).

K_JSONIQ : J S O N I Q;
K_ENCODING : E N C O D I N G;
K_VERSION : V E R S I O N;
K_MODULE : M O D U L E;
K_NAMESPACE : N A M E S P A C E;
K_DECLARE : D E C L A R E;
K_ORDERING : O R D E R I N G;
K_ORDERED: O R D E R E D;
K_UNORDERED: U N O R D E R E D;
K_COLLATION : C O L L A T I O N;
K_DEFAULT : D E F A U L T;
K_ORDER : O R D E R;
K_STABLE: S T A B L E;
K_EMPTY : E M P T Y;
K_GREATEST : G R E A T E S T;
K_LEAST : L E A S T;
K_DECIMAL_FORMAT : D E C I M A L '-' F O R M A T;
K_DECIMAL_SEPARATOR : D E C I M A L '-' S E P A R A T O R;
K_GROUPING_SEPARATOR : G R O U P I N G '-' S E P A R A T O R;
K_INFINITY : I N F I N I T Y;
K_MINUS_SIGN : M I N U S '-' S I G N;
K_NAN : N A N;
K_PERCENT : P E R C E N T;
K_PER_MILLE : P E R '-' M I L L E;
K_ZERO_DIGIT : Z E R O '-' D I G I T;
K_DIGIT : D I G I T;
K_PATTERN_SEPARATOR : P A T T E R N '-' S E P A R A T O R;
K_IMPORT : I M P O R T;
K_AT : A T;
K_VARIABLE : V A R I A B L E;
K_EXTERNAL : E X T E R N A L;
K_FUNCTION : F U N C T I O N;
K_AS : A S ;
K_RETURN : R E T U R N;
K_FOR : F O R;
K_LET: L E T;
K_ALLOWING : A L L O W I N G;
K_IN : I N;
K_OF : O F;
K_TO : T O;
K_COUNT : C O U N T;
K_WHERE : W H E R E;
K_GROUP : G R O U P;
K_BY : B Y;
K_ASCENDING : A S C E N D I N G;
K_DESCENDING : D E S C E N D I N G;
K_SOME: S O M E;
K_EVERY: E V E R Y;
K_SATISFIES: S A T I S F I E S;
K_SWITCH: S W I T C H;
K_CASE: C A S E;
K_TYPESWITCH : T Y P E S W I T C H;
K_IF: I F;
K_THEN: T H E N;
K_ELSE: E L S E;
K_TRY: T R Y;
K_CATCH: C A T C H;
K_AND: A N D;
K_OR: O R;
K_NOT: N O T;
K_EQ: E Q;
K_NE: N E;
K_LT: L T;
K_LE: L E;
K_GT: G T;
K_GE: G E;
K_DIV: D I V;
K_IDIV: I D I V;
K_MOD: M O D;
K_INSTANCE: I N S T A N C E;
K_TREAT: T R E A T;
K_CASTABLE: C A S T A B L E;
K_CAST: C A S T;
K_ITEM: I T E M;
K_OBJECT: O B J E C T;
K_ARRAY: A R R A Y;
K_JSON_ITEM: J S O N '-' I T E M;
K_ATOMIC: A T O M I C;
K_STRING: S T R I N G;
K_INTEGER: I N T E G E R;
K_DECIMAL: D E C I M A L;
K_DOUBLE: D O U B L E;
K_BOOLEAN: B O O L E A N;
K_LONG: L O N G;
K_SHORT: S H O R T;
K_BYTE: B Y T E;
K_FLOAT: F L O A T;
K_DATE: D A T E;
K_DATETIME: D A T E T I M E;
K_DATETIMESTAMP: D A T E T I M E S T A M P;
K_GDAY: G D A Y;
K_GMONTH: G M O N T H;
K_GMONTHDAY: G M O N T H D A Y;
K_GYEAR: G Y E A R;
K_GYEARMONTH: G Y E A R M O N T H;
K_TIME: T I M E;
K_DURATION: D U R A T I O N;
K_DAYTIMEDURATION: D A Y T I M E D U R A T I O N;
K_YEARMONTHDURATION: Y E A R M O N T H D U R A T I O N;
K_BASE64NIBARY: B A S E '6' '4' B I N A R Y;
K_HEXBINARY: H E X B I N A R Y;
K_ANYURI: A N Y U R I;
K_NULL: N U L L;

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

// Literals 

literal  : numericLiteral | StringLiteral | booleanLiteral | nullLiteral ;
numericLiteral : IntegerLiteral | DecimalLiteral | DoubleLiteral ;
booleanLiteral : 'true' | 'false' ;
nullLiteral : 'null' ;

IntegerLiteral : Digits ;

DecimalLiteral : ('.' Digits) | (Digits '.' [0-9]*) ;

DoubleLiteral : (('.' Digits) | (Digits ('.' [0-9]*)?)) [eE] [+-]? Digits ;

StringLiteral : '"' ( PredefinedCharRef | CharRef | ~["\\] )* '"' ;

fragment
PredefinedCharRef: '\u005C' ('\u005C'|'\u002F'|'\u0027'|'b'|'f'|'n'|'r'|'t'|'"');

fragment
CharRef : '\\' 'u' [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F][0-9a-fA-F];

Comment : '(:' (CommentContents | Comment)* ':)' ;

fragment
CommentContents: ~(':'|'\u0028'|'\u0029');
// This is XML Name without the ':' character

NCName : NameStartChar (NameChar)*;

fragment
NameStartChar: 
	[a-zA-Z] |
	'\u00C0'..'\u00D6' |
	'\u00D8'..'\u00F6' |
	'\u00F8'..'\u02FF' |
	'\u0370'..'\u037D' |
	'\u037F'..'\u1FFF' |
	'\u200C'..'\u200D' |
	'\u2070'..'\u218F' |
	'\u2C00'..'\u2FEF' |
	'\u3001'..'\uD7FF' |
	'\uF900'..'\uFDCF' |
	'\uFDF0'..'\uFFFD' ; 

fragment
NameChar:
	NameStartChar | '-' | '.' |
	[0-9] |
   	'\u00B7' |
        '\u0300'..'\u036F' |
        '\u203F'..'\u2040';

fragment
Digits   : [0-9]+ ;

WS
    : (' '
    | '\t'
    | '\n'
    | '\r')+ ->skip
    ;