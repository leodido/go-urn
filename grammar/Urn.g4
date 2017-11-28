/*
 * This grammar encodes the URN syntax.
 *
 * RFC 2141.
 * https.//tools.ietf.org/html/rfc2141
 */
grammar Urn;

@parser::members {
func isIdentifier(s string) bool {
    for i, r := range s {
        // !unicode.IsLetter(r) etc. when (if) we'll need unicode
        if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') && (r != '-' || (i == 0 && r == '-')) {
            return false
        }
    }
    return true
}
}

urn
    : Urn Colon iD Colon sS EOF
    ;

iD  
    : Part
    {len($ctx.GetText()) <= 32}?<fail={"exceed max (32) number of characters"}>
    {isIdentifier($ctx.GetText())}?<fail={"is not a valid identifier"}>
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

Whitespace
    : ' '
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