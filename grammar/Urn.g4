/*
 * This grammar encodes the URN syntax.
 *
 * RFC 2141.
 * https.//tools.ietf.org/html/rfc2141
 */
grammar Urn;

urn
    : Urn Colon iD Colon sS EOF
    ;

iD  
    : Part {len($ctx.GetText()) <= 32}?
    ;

sS
    : Part (Colon|Hyphen)* sS?
    | (Colon|Hyphen)+ sS?
    | Part sS?
    | Urn
    ;

Urn
    : URN
    ;

Part
    : IDENTIFIER
    | CHARS
    ;

Colon
    : ':'
    ;

Hyphen
    : '-'
    ;

fragment URN
    : [uU][rR][nN]
    ;

fragment IDENTIFIER   
    : ALPHA_NUMERIC (ALPHA_NUMERIC|Hyphen)+
    ;

fragment CHARS
    : (TRANSLATION|HEX)+
    ;

fragment HEX
    : '%' HEX_DIGIT HEX_DIGIT
    ;

fragment NUMBER
    : [0-9]
    ;

fragment LETTER
    : [A-Z]
    | [a-z]
    ;

fragment ALPHA_NUMERIC
    : NUMBER
    | LETTER 
    ;

fragment OTHER
    : '('
    | ')'
    | '+'
    | ','
    | '.'
    | '='
    | '@'
    | ';'
    | '$'
    | '_'
    | '!'
    | '*'
    | '\''
    ;

fragment HEX_DIGIT
    : NUMBER
    | [A-F]
    | [a-f]
    ;

fragment RESERVED
    : '%'
    | '/'
    | '?'
    | '#'
    ;

fragment TRANSLATION
    : OTHER
    | RESERVED
    | ALPHA_NUMERIC
    ;