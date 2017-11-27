/*
 * This grammar encodes the URN syntax.
 *
 * RFC 2141: https.//tools.ietf.org/html/rfc2141
 */
grammar Urn;

@parser::header {
}

@parser::members {
}

urn
    : Urn Colon iD Colon sS EOF
    ;

iD  
    : Part {len($ctx.GetText()) <= 32}?
    ;

sS
    : Part (Colon Part)*
    | Urn
    ;

Urn
    : URN
    ;

Part
    : IDENTIFIER|CHARS
    ;

Colon
    : ':'
    ;

fragment URN
    : [uU][rR][nN]
    ;

fragment IDENTIFIER   
    : ALPHA_NUMERIC (ALPHA_NUMERIC|HYPHEN)+
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

fragment HYPHEN
    : '-'
    ;

fragment OTHER
    : '('
    | ')'
    | '+'
    | ','
    | HYPHEN
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
    : ALPHA_NUMERIC
    | OTHER
    | RESERVED
    ;