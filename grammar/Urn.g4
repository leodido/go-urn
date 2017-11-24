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
    : 'urn' Colon iD Colon sS EOF
    ;

iD
    : 'NID'
    ;

sS
    : 'NSS'
    ;

Colon
    : ':'
    ;