package urn

import (
	"fmt"
	"strings"

	pcre "github.com/gijsbers/go-pcre"
)

var re = `^
	(?:(?P<pre>[uU][rR][nN]):(?!urn:))

	(?P<nid>[A-Za-z0-9][A-Za-z0-9-]{0,31}):

	(?P<nss>(?:[A-Za-z0-9()+,\-.:=@;$_!*']|[%][A-Fa-f0-9][A-Fa-f0-9])+)
$`

var pattern = pcre.MustCompile(re, pcre.EXTENDED)
var hexrepr = pcre.MustCompile("[%][a-f0-9]{2}", pcre.CASELESS)

// URN represents an Uniform Resource Name.
//
// The general form represented is:
//
//	urn:<id>:<ss>
//
// Details at https://tools.ietf.org/html/rfc2141
type URN struct {
	ID string // Namespace identifier
	SS string // Namespace specific string
}

// Parse is ...
func Parse(u string) (*URN, bool) {
	matcher := pattern.MatcherString(u, 0)
	matches := matcher.Matches()

	if matches {
		urn := &URN{}
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
func (u *URN) String() string {
	var res string
	if u.ID != "" && u.SS != "" {
		res = "urn:" + u.ID + ":" + u.SS
	}

	return res
}

// Normalize is ...
func (u *URN) Normalize() *URN {
	matcher := hexrepr.MatcherString(u.SS, 0)
	results := matcher.ExtractString()

	fmt.Println("results> ", results)
	// find all hex within u.SS
	// lowercase any match
	// lowercase u.ID
	// reconstruct string

	return &URN{
		ID: strings.ToLower(u.ID),
		// SS: ...
	}
}

/*
func matchAll(re pcre.Regexp, subject []byte, flags int) [][]byte {
	m := re.Matcher(subject, 0)
	all := [][]byte{}
	for m.Match(subject, flags) {
		all = append(all, subject[m.ovector[0]:m.ovector[1]])
		subject = subject[m.ovector[1]:]
	}
	return all
}
*/

// Equal is ...
func (u *URN) Equal(x *URN) bool {
	return *u.Normalize() == *x.Normalize()
}
