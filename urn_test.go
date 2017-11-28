package urn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	in   string // the input
	ok   bool   // whether it is valid or not
	obj  *URN   // a pointer to the resulting urn.URN instance
	col  int    // the colum where the parsing error occurres
	tree string // the tree representation
}{
	// ok
	{"urn:simple:simple", true, &URN{ID: "simple", SS: "simple"}, -1, "(urn urn : (iD simple) : (sS simple) <EOF>)"},

	// ok - RFC examples
	{"URN:foo:a123,456", true, &URN{ID: "foo", SS: "a123,456"}, -1, "(urn URN : (iD foo) : (sS a123,456) <EOF>)"},
	{"urn:foo:a123,456", true, &URN{ID: "foo", SS: "a123,456"}, -1, "(urn urn : (iD foo) : (sS a123,456) <EOF>)"},
	{"urn:FOO:a123,456", true, &URN{ID: "FOO", SS: "a123,456"}, -1, "(urn urn : (iD FOO) : (sS a123,456) <EOF>)"},
	{"urn:foo:A123,456", true, &URN{ID: "foo", SS: "A123,456"}, -1, "(urn urn : (iD foo) : (sS A123,456) <EOF>)"},
	{"urn:foo:a123%2C456", true, &URN{ID: "foo", SS: "a123%2C456"}, -1, "(urn urn : (iD foo) : (sS a123%2C456) <EOF>)"},
	{"URN:FOO:a123%2c456", true, &URN{ID: "FOO", SS: "a123%2c456"}, -1, "(urn URN : (iD FOO) : (sS a123%2c456) <EOF>)"},

	// ok - SCIM v2
	{"urn:ietf:params:scim:schemas:core:2.0:User", true, &URN{ID: "ietf", SS: "params:scim:schemas:core:2.0:User"}, -1, "(urn urn : (iD ietf) : (sS params : scim : schemas : core : 2.0 : User) <EOF>)"},
	{"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", true, &URN{ID: "ietf", SS: "params:scim:schemas:extension:enterprise:2.0:User"}, -1, "(urn urn : (iD ietf) : (sS params : scim : schemas : extension : enterprise : 2.0 : User) <EOF>)"},

	// ok - minimum urn
	{"urn:a:b", true, &URN{ID: "a", SS: "b"}, -1, "(urn urn : (iD a) : (sS b) <EOF>)"},

	// ok - URN prefix is case-insensitive
	{"URN:simple:simple", true, &URN{ID: "simple", SS: "simple"}, -1, "(urn URN : (iD simple) : (sS simple) <EOF>)"},
	{"Urn:simple:simple", true, &URN{ID: "simple", SS: "simple"}, -1, "(urn Urn : (iD simple) : (sS simple) <EOF>)"},

	// ok - ID can contain the "urn" string but it can not be exactly equal to it
	{"urn:urna:simple", true, &URN{ID: "urna", SS: "simple"}, -1, "(urn urn : (iD urna) : (sS simple) <EOF>)"},
	{"urn:burnout:nss", true, &URN{ID: "burnout", SS: "nss"}, -1, "(urn urn : (iD burnout) : (sS nss) <EOF>)"},
	{"urn:burn:nss", true, &URN{ID: "burn", SS: "nss"}, -1, "(urn urn : (iD burn) : (sS nss) <EOF>)"},
	{"urn:urnurnurn:x", true, &URN{ID: "urnurnurn", SS: "x"}, -1, "(urn urn : (iD urnurnurn) : (sS x) <EOF>)"},

	// ok - ID can contains maximum 32 characters
	{"urn:abcdefghilmnopqrstuvzabcdefghilm:x", true, &URN{ID: "abcdefghilmnopqrstuvzabcdefghilm", SS: "x"}, 4 + 32, "(urn urn : (iD abcdefghilmnopqrstuvzabcdefghilm) : (sS x) <EOF>)"},

	// ok - ID can be alpha numeric
	{"URN:123:x", true, &URN{ID: "123", SS: "x"}, -1, "(urn URN : (iD 123) : (sS x) <EOF>)"},
	{"URN:1ab:x", true, &URN{ID: "1ab", SS: "x"}, -1, "(urn URN : (iD 1ab) : (sS x) <EOF>)"},
	{"URN:a1b:x", true, &URN{ID: "a1b", SS: "x"}, -1, "(urn URN : (iD a1b) : (sS x) <EOF>)"},
	{"URN:a12:x", true, &URN{ID: "a12", SS: "x"}, -1, "(urn URN : (iD a12) : (sS x) <EOF>)"},
	{"URN:cd2:x", true, &URN{ID: "cd2", SS: "x"}, -1, "(urn URN : (iD cd2) : (sS x) <EOF>)"},

	// ok - ID can contain an hyphen but not in its first position
	{"URN:abcd-:x", true, &URN{ID: "abcd-", SS: "x"}, -1, "(urn URN : (iD abcd-) : (sS x) <EOF>)"},
	{"URN:abcd-abcd:x", true, &URN{ID: "abcd-abcd", SS: "x"}, -1, "(urn URN : (iD abcd-abcd) : (sS x) <EOF>)"},
	{"URN:a123-456z:x", true, &URN{ID: "a123-456z", SS: "x"}, -1, "(urn URN : (iD a123-456z) : (sS x) <EOF>)"},

	// ok - SS can exactly contain the "urn" string
	{"urn:urnx:urn", true, &URN{ID: "urnx", SS: "urn"}, -1, "(urn urn : (iD urnx) : (sS urn) <EOF>)"},
	{"urn:urnurnurn:urn", true, &URN{ID: "urnurnurn", SS: "urn"}, -1, "(urn urn : (iD urnurnurn) : (sS urn) <EOF>)"},

	// ok - SS can contains multiple colons
	{"urn:ciao:a:b:c", true, &URN{ID: "ciao", SS: "a:b:c"}, -1, "(urn urn : (iD ciao) : (sS a : b : c) <EOF>)"},

	// no - ID can not start with an hyphen
	// {"URN:-xxx:x", false, nil, 5, ""}, // (fixme) - hyphen probably causes overlap between IDENTIFIER and CHARS

	// no - ID can not contains more than 32 characters
	{"urn:abcdefghilmnopqrstuvzabcdefghilmn:specificstring", false, nil, 4 + 33, ""},

	// no - ID can not be equal to "urn"
	{"urn:urn:NSS", false, nil, 4, ""},

	// no - Incomplete URNs
	{"urn:", false, nil, 4, ""},
	{"urn:a", false, nil, 5, ""},
	{"urn:a:", false, nil, 6, ""},
	// {"urn:a:x:", false, nil, 8, ""}, // (fixme) - probable issue with last colon
}

func TestUrnParse(t *testing.T) {
	for _, tt := range tests {
		urn, err := Parse(tt.in)

		if ok := err == nil; ok {
			require.True(t, tt.ok)
			// Ignoring column testing since there is no error
			assert.Equal(t, tt.obj.ID, urn.ID)
			assert.Equal(t, tt.obj.SS, urn.SS)
			assert.Equal(t, tt.tree, urn.Tree())
		} else {
			require.False(t, tt.ok)
			assert.Equal(t, tt.col, err.(*Error).Column)
			assert.Empty(t, urn)
		}
	}
}
