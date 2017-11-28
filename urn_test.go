package urn

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	in   string // the input
	ok   bool   // whether it is valid or not
	obj  *URN   // a pointer to the resulting urn.URN instance
	col  int    // the colum where the parsing error occurres
	tree string // the tree representation
}

var tests = []testCase{
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
	{"urn:a::", true, &URN{ID: "a", SS: ":"}, -1, "(urn urn : (iD a) : (sS :) <EOF>)"},
	{"urn:a:-", true, &URN{ID: "a", SS: "-"}, -1, "(urn urn : (iD a) : (sS -) <EOF>)"},
	{"urn:a:%", true, &URN{ID: "a", SS: "%"}, -1, "(urn urn : (iD a) : (sS %) <EOF>)"},

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

	// ok - ID can contain an hyphen (not in its first position, see below)
	{"URN:abcd-:x", true, &URN{ID: "abcd-", SS: "x"}, -1, "(urn URN : (iD abcd-) : (sS x) <EOF>)"},
	{"URN:abcd-abcd:x", true, &URN{ID: "abcd-abcd", SS: "x"}, -1, "(urn URN : (iD abcd-abcd) : (sS x) <EOF>)"},
	{"URN:a123-456z:x", true, &URN{ID: "a123-456z", SS: "x"}, -1, "(urn URN : (iD a123-456z) : (sS x) <EOF>)"},

	// ok - SS can contain the "urn" string, also be exactly equal to it
	{"urn:urnx:urn", true, &URN{ID: "urnx", SS: "urn"}, -1, "(urn urn : (iD urnx) : (sS urn) <EOF>)"},
	{"urn:urnurnurn:urn", true, &URN{ID: "urnurnurn", SS: "urn"}, -1, "(urn urn : (iD urnurnurn) : (sS urn) <EOF>)"},
	{"urn:hey:urnurnurn", true, &URN{ID: "hey", SS: "urnurnurn"}, -1, "(urn urn : (iD hey) : (sS urnurnurn) <EOF>)"},

	// ok - SS can contains and discerns multiple colons, also at the end
	{"urn:ciao:a:b:c", true, &URN{ID: "ciao", SS: "a:b:c"}, -1, "(urn urn : (iD ciao) : (sS a : b : c) <EOF>)"},
	{"urn:aaa:x:y:", true, &URN{ID: "aaa", SS: "x:y:"}, -1, "(todo)"},
	{"urn:aaa:x:y:", true, &URN{ID: "aaa", SS: "x:y:"}, -1, "(todo)"},

	// ok - SS can contain (and also start with) some non-alphabetical characters
	{"urn:ciao:-", true, &URN{ID: "ciao", SS: "-"}, -1, "(urn urn : (iD ciao) : (sS -) <EOF>)"},
	{"urn:ciao::", true, &URN{ID: "ciao", SS: ":"}, -1, "(urn urn : (iD ciao) : (sS :) <EOF>)"},
	{"urn:ciao:!", true, &URN{ID: "ciao", SS: "!"}, -1, "(urn urn : (iD ciao) : (sS !) <EOF>)"},
	{"urn:ciao:!?", true, &URN{ID: "ciao", SS: "!?"}, -1, "(urn urn : (iD ciao) : (sS !?) <EOF>)"},
	{"urn:ciao:-!:?-,:x", true, &URN{ID: "ciao", SS: "-!:?-,:x"}, -1, "(urn urn : (iD ciao) : (sS -!:?-,:x) <EOF>)"},
	{"urn:ciao:###", true, &URN{ID: "ciao", SS: "###"}, -1, "(urn urn : (iD ciao) : (sS ###) <EOF>)"},
	{"urn:ciao:#?!#(xyz)+a,b.*@g=$_'", true, &URN{ID: "ciao", SS: "#?!#(xyz)+a,b.*@g=$_'"}, -1, "(urn urn : (iD ciao) : (sS #?!#(xyz)+a,b.*@g=$_') <EOF>)"},

	// ok - SS can contain (and also start with) hexadecimal representation of octets  // (todo)

	// no - ID can not start with an hyphen
	{"URN:-xxx:x", false, nil, 4, ""},
	{"URN:---xxx:x", false, nil, 4, ""},

	// no - ID can not start with a colon
	{"urn::colon:nss", false, nil, 4, ""},
	{"urn::::nss", false, nil, 4, ""},

	// no - ID can not contains more than 32 characters
	{"urn:abcdefghilmnopqrstuvzabcdefghilmn:specificstring", false, nil, 4 + 33, ""},

	// no - ID can not contain special characters
	{"URN:a!?:x", false, nil, 7, ""},
	{"URN:#,:x", false, nil, 6, ""},
	{"URN:bc'.@:x", false, nil, 9, ""},

	// no - ID can not be equal to "urn"
	{"urn:urn:NSS", false, nil, 4, ""},

	// no - ID can not contain spaces
	{"urn:white space:NSS", false, nil, 9, ""},

	// no - SS can not contain spaces
	{"urn:concat:no spaces", false, nil, 13, ""},

	// no - Incomplete URNs
	{"urn:", false, nil, 4, ""},
	{"urn::", false, nil, 4, ""},
	{"urn:a", false, nil, 5, ""},
	{"urn:a:", false, nil, 6, ""},
}

func herror(index int, test testCase) string {
	return "Test case num. " + strconv.Itoa(index+1) + ", input \"" + test.in + "\""
}

func TestUrnParse(t *testing.T) {
	for ii, tt := range tests {
		urn, err := Parse(tt.in)

		if ok := err == nil; ok {
			require.True(t, tt.ok, herror(ii, tt))
			require.Equal(t, tt.obj.ID, urn.ID, herror(ii, tt))
			require.Equal(t, tt.obj.SS, urn.SS, herror(ii, tt))
			// require.Equal(t, tt.tree, urn.Tree(), herror(ii, tt)) // (fixme) > flatten tree?
			// Ignoring column testing since there is no error
		} else {
			require.False(t, tt.ok, herror(ii, tt))
			require.Equal(t, tt.col, err.(*Error).Column, herror(ii, tt))
			require.Empty(t, urn, herror(ii, tt))
			//require.Empty(t, urn.Tree(), herror(ii, tt))
		}
	}
}
