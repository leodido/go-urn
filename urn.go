// Package urn implements RFC 2141.
//
// It provides a parser capable, given a string respecting valid syntax as per RFC 2141, to instantiate an URN type.
// Furthermore this package provides methods to perform normalization and lexical equivalence on URN instances.
package urn

import (
	"strings"

	pcre "github.com/gijsbers/go-pcre"
)

var re = `^
	(?:(?P<pre>[uU][rR][nN]):(?!urn:))

	(?P<nid>[A-Za-z0-9][A-Za-z0-9-]{0,31}):

	(?P<nss>(?:[A-Za-z0-9()+,\-.:=@;$_!*']|[%][A-Fa-f0-9][A-Fa-f0-9])+)
$`

var pattern = pcre.MustCompile(re, pcre.EXTENDED)
var hexrepr = pcre.MustCompile("[%][A-F0-9]{2}", 0)

// URN represents an Uniform Resource Name.
//
// The general form represented is:
//
//	urn:<id>:<ss>
//
// Details at https://tools.ietf.org/html/rfc2141.
type URN struct {
	prefix string // Static prefix. Equal to "urn" when empty.
	ID     string // Namespace identifier
	SS     string // Namespace specific string
}

// Parse is responsible to create an URN instance from a string matching the correct URN syntax.
func Parse(u string) (*URN, bool) {
	matcher := pattern.MatcherString(u, 0)
	matches := matcher.Matches()

	if matches {
		urn := &URN{}
		urn.prefix, _ = matcher.NamedString("pre")
		urn.ID, _ = matcher.NamedString("nid")
		urn.SS, _ = matcher.NamedString("nss")

		return urn, matches
	}

	return nil, matches
}

// String reassembles the URN into a valid URN string.
//
// This requires both ID and SS fields to be non-empty.
// Otherwise it returns an empty string.
//
// Default URN prefix is "urn".
func (u *URN) String() string {
	var res string
	if u.ID != "" && u.SS != "" {
		if u.prefix == "" {
			res += "urn"
		}
		res += u.prefix + ":" + u.ID + ":" + u.SS
	}

	return res
}

// Normalize turns the receiving URN into its norm version.
//
// Which means: lowercase prefix, lowercase namespace identifier, and immutate namespace specific string chars (except <hex> tokens which are lowercased).
func (u *URN) Normalize() *URN {
	norm := ""
	ss := u.SS
	matcher := hexrepr.MatcherString(ss, 0)
	for matcher.MatchString(ss, 0) {
		indexes := matcher.Index()
		from := indexes[0]
		to := indexes[1]
		norm += ss[:from] + strings.ToLower(ss[from:to])
		ss = ss[to:]
	}
	norm += ss

	return &URN{
		prefix: "urn",
		ID:     strings.ToLower(u.ID),
		SS:     norm,
	}
}

// Equal checks the lexical equivalence of the current URN with another one.
func (u *URN) Equal(x *URN) bool {
	return *u.Normalize() == *x.Normalize()
}
