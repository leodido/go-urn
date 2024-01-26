package urn

import (
	"fmt"
	"strconv"
	"strings"
)

func ierror(index int) string {
	return "Test case num. " + strconv.Itoa(index+1)
}

func herror(index int, test testCase) string {
	return ierror(index) + ", input \"" + string(test.in) + "\""
}

func rxpad(str string, lim int) string {
	str = str + strings.Repeat(" ", lim)
	return str[:lim]
}

type testCase struct {
	in     []byte // the input
	ok     bool   // whether it is valid or not
	obj    *URN   // a pointer to the resulting urn.URN instance
	str    string // string representation
	norm   string // norm string representation
	estr   string // error string
	isSCIM bool   // whether it is a SCIM URN or not
}

var urnlexTestCases = []testCase{
	// Italian act
	{
		[]byte("urn:lex:it:stato:legge:2003-09-21;456"),
		true,
		&URN{
			prefix: "urn",
			ID:     "lex",
			SS:     "it:stato:legge:2003-09-21;456",
		},
		"urn:lex:it:stato:legge:2003-09-21;456",
		"urn:lex:it:stato:legge:2003-09-21;456",
		"",
		false,
	},
	// Italian decree
	// fixme(leodido)
	// verify whether it is correct or not that ~ is not accepted
	// does it requires RFC 8141 (issue #17) ?
	// {
	// 	[]byte("urn:lex:it:ministero.giustizia:decreto:1992-07-24;358~art5"),
	// 	true,
	// 	&URN{
	// 		prefix: "urn",
	// 		ID:     "lex",
	// 		SS:     "it:ministero.giustizia:decreto:1992-07-24;358~art5",
	// 	},
	// 	"it:ministero.giustizia:decreto:1992-07-24;358~art5",
	// 	"it:ministero.giustizia:decreto:1992-07-24;358~art5",
	// 	"",
	// 	false,
	// },
	// French act
	{
		[]byte("urn:lex:fr:etat:lois:2004-12-06;321"),
		true,
		&URN{
			prefix: "urn",
			ID:     "lex",
			SS:     "fr:etat:lois:2004-12-06;321",
		},
		"urn:lex:fr:etat:lois:2004-12-06;321",
		"urn:lex:fr:etat:lois:2004-12-06;321",
		"",
		false,
	},
	// Spanish act
	{
		[]byte("urn:lex:es:estado:ley:2002-07-12;123"),
		true,
		&URN{
			prefix: "urn",
			ID:     "lex",
			SS:     "es:estado:ley:2002-07-12;123",
		},
		"urn:lex:es:estado:ley:2002-07-12;123",
		"urn:lex:es:estado:ley:2002-07-12;123",
		"",
		false,
	},
	// Glarus Swiss Canton decree
	{
		[]byte("urn:lex:ch;glarus:regiere:erlass:2007-10-15;963"),
		true,
		&URN{
			prefix: "urn",
			ID:     "lex",
			SS:     "ch;glarus:regiere:erlass:2007-10-15;963",
		},
		"urn:lex:ch;glarus:regiere:erlass:2007-10-15;963",
		"urn:lex:ch;glarus:regiere:erlass:2007-10-15;963",
		"",
		false,
	},
	// EU Council Directive
	{
		[]byte("urn:lex:eu:council:directive:2010-03-09;2010-19-UE"),
		true,
		&URN{
			prefix: "urn",
			ID:     "lex",
			SS:     "eu:council:directive:2010-03-09;2010-19-UE",
		},
		"urn:lex:eu:council:directive:2010-03-09;2010-19-UE",
		"urn:lex:eu:council:directive:2010-03-09;2010-19-UE",
		"",
		false,
	},
	{
		[]byte("urn:lex:eu:council:directive:2010-03-09;2010-19-UE"),
		true,
		&URN{
			prefix: "urn",
			ID:     "lex",
			SS:     "eu:council:directive:2010-03-09;2010-19-UE",
		},
		"urn:lex:eu:council:directive:2010-03-09;2010-19-UE",
		"urn:lex:eu:council:directive:2010-03-09;2010-19-UE",
		"",
		false,
	},
	// US FSC decision
	{
		[]byte("urn:lex:us:federal.supreme.court:decision:1963-03-18;372.us.335"),
		true,
		&URN{
			prefix: "urn",
			ID:     "lex",
			SS:     "us:federal.supreme.court:decision:1963-03-18;372.us.335",
		},
		"urn:lex:us:federal.supreme.court:decision:1963-03-18;372.us.335",
		"urn:lex:us:federal.supreme.court:decision:1963-03-18;372.us.335",
		"",
		false,
	},
}

var scimOnlyTestCases = []testCase{
	// ok
	{
		[]byte("urn:ietf:params:scim:schemas:core:2.0:User"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "schemas:core:2.0:User",
		},
		"urn:ietf:params:scim:schemas:core:2.0:User",
		"urn:ietf:params:scim:schemas:core:2.0:User",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "schemas:extension:enterprise:2.0:User",
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "schemas:extension:enterprise:2.0:User:userName",
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName",
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:api:messages:2.0:ListResponse"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "api:messages:2.0:ListResponse",
		},
		"urn:ietf:params:scim:api:messages:2.0:ListResponse",
		"urn:ietf:params:scim:api:messages:2.0:ListResponse",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:core"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "schemas:core",
		},
		"urn:ietf:params:scim:schemas:core",
		"urn:ietf:params:scim:schemas:core",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:param:core"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "param:core",
		},
		"urn:ietf:params:scim:param:core",
		"urn:ietf:params:scim:param:core",
		"",
		true,
	},

	// no
	{
		[]byte("arn:ietf:params:scim:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 0),
		false,
	},
	{
		[]byte("usn:ietf:params:scim:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 1),
		false,
	},
	{
		[]byte("urm:ietf:params:scim:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 2),
		false,
	},
	{
		[]byte("urno:ietf:params:scim:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 3),
		false,
	},
	{
		[]byte("urno"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 3),
		false,
	},
	{
		[]byte("urn:WRONG:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errSCIMNamespace, 4),
		false,
	},
	{
		[]byte("urn:ietf:params:scim:WRONG:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errSCIMType, 21),
		false,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:$"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errSCIMName, 29),
		false,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:core-"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errSCIMName, 33),
		false,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:core:"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errSCIMOtherIncomplete, 33),
		false,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:core:2.&"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errSCIMOther, 36),
		false,
	},
}

var urn2141OnlyTestCases = []testCase{
	// ok
	{
		[]byte("urn:simple:simple"),
		true,
		&URN{
			prefix: "urn",
			ID:     "simple",
			SS:     "simple",
		},
		"urn:simple:simple",
		"urn:simple:simple",
		"",
		false,
	},
	{
		[]byte("urn:ciao:%5D"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     "%5D",
		},
		"urn:ciao:%5D",
		"urn:ciao:%5d",
		"",
		false,
	},

	// ok - RFC examples
	{
		[]byte("URN:foo:a123,456"),
		true,
		&URN{
			prefix: "URN",
			ID:     "foo",
			SS:     "a123,456",
		},
		"URN:foo:a123,456",
		"urn:foo:a123,456",
		"",
		false,
	},
	{
		[]byte("urn:foo:a123,456"),
		true,
		&URN{
			prefix: "urn",
			ID:     "foo",
			SS:     "a123,456",
		},
		"urn:foo:a123,456",
		"urn:foo:a123,456",
		"",
		false,
	},
	{
		[]byte("urn:FOO:a123,456"),
		true,
		&URN{
			prefix: "urn",
			ID:     "FOO",
			SS:     "a123,456",
		},
		"urn:FOO:a123,456",
		"urn:foo:a123,456",
		"",
		false,
	},
	{
		[]byte("urn:foo:A123,456"),
		true,
		&URN{
			prefix: "urn",
			ID:     "foo",
			SS:     "A123,456",
		},
		"urn:foo:A123,456",
		"urn:foo:A123,456",
		"",
		false,
	},
	{
		[]byte("urn:foo:a123%2C456"),
		true,
		&URN{
			prefix: "urn",
			ID:     "foo",
			SS:     "a123%2C456",
		},
		"urn:foo:a123%2C456",
		"urn:foo:a123%2c456",
		"",
		false,
	},
	{
		[]byte("URN:FOO:a123%2c456"),
		true,
		&URN{
			prefix: "URN",
			ID:     "FOO",
			SS:     "a123%2c456",
		},
		"URN:FOO:a123%2c456",
		"urn:foo:a123%2c456",
		"",
		false,
	},
	{
		[]byte("URN:FOO:ABC%FFabc123%2c456"),
		true,
		&URN{
			prefix: "URN",
			ID:     "FOO",
			SS:     "ABC%FFabc123%2c456",
		},
		"URN:FOO:ABC%FFabc123%2c456",
		"urn:foo:ABC%ffabc123%2c456",
		"",
		false,
	},
	{
		[]byte("URN:FOO:ABC%FFabc123%2C456%9A"),
		true,
		&URN{
			prefix: "URN",
			ID:     "FOO",
			SS:     "ABC%FFabc123%2C456%9A",
		},
		"URN:FOO:ABC%FFabc123%2C456%9A",
		"urn:foo:ABC%ffabc123%2c456%9a",
		"",
		false,
	},

	// ok - SCIM v2
	{
		[]byte("urn:ietf:params:scim:schemas:core:2.0:User"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf",
			SS:     "params:scim:schemas:core:2.0:User",
		},
		"urn:ietf:params:scim:schemas:core:2.0:User",
		"urn:ietf:params:scim:schemas:core:2.0:User",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf",
			SS:     "params:scim:schemas:extension:enterprise:2.0:User",
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf",
			SS:     "params:scim:schemas:extension:enterprise:2.0:User:userName",
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName",
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf",
			SS:     "params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified",
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified",
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified",
		"",
		true,
	},

	// ok - minimum urn
	{
		[]byte("urn:a:b"),
		true,
		&URN{
			prefix: "urn",
			ID:     "a",
			SS:     "b",
		},
		"urn:a:b",
		"urn:a:b",
		"",
		false,
	},
	{
		[]byte("urn:a::"),
		true,
		&URN{
			prefix: "urn",
			ID:     "a",
			SS:     ":",
		},
		"urn:a::",
		"urn:a::",
		"",
		false,
	},
	{
		[]byte("urn:a:-"),
		true,
		&URN{
			prefix: "urn",
			ID:     "a",
			SS:     "-",
		},
		"urn:a:-",
		"urn:a:-",
		"",
		false,
	},

	// ok - URN prefix is case-insensitive
	{
		[]byte("URN:simple:simple"),
		true,
		&URN{
			prefix: "URN",
			ID:     "simple",
			SS:     "simple",
		},
		"URN:simple:simple",
		"urn:simple:simple",
		"",
		false,
	},
	{
		[]byte("Urn:simple:simple"),
		true,
		&URN{
			prefix: "Urn",
			ID:     "simple",
			SS:     "simple",
		},
		"Urn:simple:simple",
		"urn:simple:simple",
		"",
		false,
	},

	// ok - ID can contain the "urn" string but it can not be exactly equal to it
	{
		[]byte("urn:urna:simple"),
		true,
		&URN{
			prefix: "urn",
			ID:     "urna",
			SS:     "simple",
		},
		"urn:urna:simple",
		"urn:urna:simple",
		"",
		false,
	},
	{
		[]byte("urn:burnout:nss"),
		true,
		&URN{
			prefix: "urn",
			ID:     "burnout",
			SS:     "nss",
		},
		"urn:burnout:nss",
		"urn:burnout:nss",
		"",
		false,
	},
	{
		[]byte("urn:burn:nss"),
		true,
		&URN{
			prefix: "urn",
			ID:     "burn",
			SS:     "nss",
		},
		"urn:burn:nss",
		"urn:burn:nss",
		"",
		false,
	},
	{
		[]byte("urn:urnurnurn:x"),
		true,
		&URN{
			prefix: "urn",
			ID:     "urnurnurn",
			SS:     "x",
		},
		"urn:urnurnurn:x",
		"urn:urnurnurn:x",
		"",
		false,
	},

	// ok - ID can contains maximum 32 characters
	{
		[]byte("urn:abcdefghilmnopqrstuvzabcdefghilm:x"),
		true,
		&URN{
			prefix: "urn",
			ID:     "abcdefghilmnopqrstuvzabcdefghilm",
			SS:     "x",
		},
		"urn:abcdefghilmnopqrstuvzabcdefghilm:x",
		"urn:abcdefghilmnopqrstuvzabcdefghilm:x",
		"",
		false,
	},

	// ok - ID can be alpha numeric
	{
		[]byte("URN:123:x"),
		true,
		&URN{
			prefix: "URN",
			ID:     "123",
			SS:     "x",
		},
		"URN:123:x",
		"urn:123:x",
		"",
		false,
	},
	{
		[]byte("URN:1ab:x"),
		true,
		&URN{
			prefix: "URN",
			ID:     "1ab",
			SS:     "x",
		},
		"URN:1ab:x",
		"urn:1ab:x",
		"",
		false,
	},
	{
		[]byte("URN:a1b:x"),
		true,
		&URN{
			prefix: "URN",
			ID:     "a1b",
			SS:     "x",
		},
		"URN:a1b:x",
		"urn:a1b:x",
		"",
		false,
	},
	{
		[]byte("URN:a12:x"),
		true,
		&URN{
			prefix: "URN",
			ID:     "a12",
			SS:     "x",
		},
		"URN:a12:x",
		"urn:a12:x",
		"",
		false,
	},
	{
		[]byte("URN:cd2:x"),
		true,
		&URN{
			prefix: "URN",
			ID:     "cd2",
			SS:     "x",
		},
		"URN:cd2:x",
		"urn:cd2:x",
		"",
		false,
	},

	// ok - ID can contain an hyphen (not in its first position, see below)
	{
		[]byte("URN:abcd-:x"),
		true,
		&URN{
			prefix: "URN",
			ID:     "abcd-",
			SS:     "x",
		},
		"URN:abcd-:x",
		"urn:abcd-:x",
		"",
		false,
	},
	{
		[]byte("URN:abcd-abcd:x"),
		true,
		&URN{
			prefix: "URN",
			ID:     "abcd-abcd",
			SS:     "x",
		},
		"URN:abcd-abcd:x",
		"urn:abcd-abcd:x",
		"",
		false,
	},
	{
		[]byte("URN:a123-456z:x"),
		true,
		&URN{
			prefix: "URN",
			ID:     "a123-456z",
			SS:     "x",
		},
		"URN:a123-456z:x",
		"urn:a123-456z:x",
		"",
		false,
	},

	// ok - SS can contain the "urn" string, also be exactly equal to it
	{
		[]byte("urn:urnx:urn"),
		true,
		&URN{
			prefix: "urn",
			ID:     "urnx",
			SS:     "urn",
		},
		"urn:urnx:urn",
		"urn:urnx:urn",
		"",
		false,
	},
	{
		[]byte("urn:urnurnurn:urn"),
		true,
		&URN{
			prefix: "urn",
			ID:     "urnurnurn",
			SS:     "urn",
		},
		"urn:urnurnurn:urn",
		"urn:urnurnurn:urn",
		"",
		false,
	},
	{
		[]byte("urn:hey:urnurnurn"),
		true,
		&URN{
			prefix: "urn",
			ID:     "hey",
			SS:     "urnurnurn",
		},
		"urn:hey:urnurnurn",
		"urn:hey:urnurnurn",
		"",
		false,
	},

	// ok - SS can contains and discerns multiple colons, also at the end
	{
		[]byte("urn:ciao:a:b:c"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     "a:b:c",
		},
		"urn:ciao:a:b:c",
		"urn:ciao:a:b:c",
		"",
		false,
	},
	{
		[]byte("urn:aaa:x:y:"),
		true,
		&URN{
			prefix: "urn",
			ID:     "aaa",
			SS:     "x:y:",
		},
		"urn:aaa:x:y:",
		"urn:aaa:x:y:",
		"",
		false,
	},
	{
		[]byte("urn:aaa:x:y:"),
		true,
		&URN{
			prefix: "urn",
			ID:     "aaa",
			SS:     "x:y:",
		},
		"urn:aaa:x:y:",
		"urn:aaa:x:y:",
		"",
		false,
	},

	// ok - SS can contain (and also start with) some non-alphabetical (ie., OTHER) characters
	{
		[]byte("urn:ciao:-"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     "-",
		},
		"urn:ciao:-",
		"urn:ciao:-",
		"",
		false,
	},
	{
		[]byte("urn:ciao::"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     ":",
		},
		"urn:ciao::",
		"urn:ciao::",
		"",
		false,
	},
	{
		[]byte("urn:colon:::::nss"),
		true,
		&URN{
			prefix: "urn",
			ID:     "colon",
			SS:     "::::nss",
		},
		"urn:colon:::::nss",
		"urn:colon:::::nss",
		"",
		false,
	},
	{
		[]byte("urn:ciao:!"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     "!",
		},
		"urn:ciao:!",
		"urn:ciao:!",
		"",
		false,
	},
	{
		[]byte("urn:ciao:!!*"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     "!!*",
		},
		"urn:ciao:!!*",
		"urn:ciao:!!*",
		"",
		false,
	},
	{
		[]byte("urn:ciao:-!:-,:x"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     "-!:-,:x",
		},
		"urn:ciao:-!:-,:x",
		"urn:ciao:-!:-,:x",
		"",
		false,
	},
	{
		[]byte("urn:ciao:=@"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     "=@",
		},
		"urn:ciao:=@",
		"urn:ciao:=@",
		"",
		false,
	},
	{
		[]byte("urn:ciao:@!=%2C(xyz)+a,b.*@g=$_'"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ciao",
			SS:     "@!=%2C(xyz)+a,b.*@g=$_'",
		},
		"urn:ciao:@!=%2C(xyz)+a,b.*@g=$_'",
		"urn:ciao:@!=%2c(xyz)+a,b.*@g=$_'",
		"",
		false,
	},

	// ok - SS can contain (and also start with) hexadecimal representation of octets
	{
		[]byte("URN:hexes:%25"),
		true,
		&URN{
			prefix: "URN",
			ID:     "hexes",
			SS:     "%25",
		},
		"URN:hexes:%25",
		"urn:hexes:%25",
		"",
		false,
	}, // Literal use of the "%" character in a namespace must be encoded using "%25"
	{
		[]byte("URN:x:abc%1Dz%2F%3az"),
		true,
		&URN{
			prefix: "URN",
			ID:     "x",
			SS:     "abc%1Dz%2F%3az",
		},
		"URN:x:abc%1Dz%2F%3az",
		"urn:x:abc%1dz%2f%3az",
		"",
		false,
	}, // Literal use of the "%" character in a namespace must be encoded using "%25"

	// no - ID can not start with an hyphen
	{
		[]byte("URN:-xxx:x"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 4]`,
		false,
	},
	{
		[]byte("URN:---xxx:x"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 4]`,
		false,
	},

	// no - ID can not start with a colon
	{
		[]byte("urn::colon:nss"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 4]`,
		false,
	},
	{
		[]byte("urn::::nss"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 4]`,
		false,
	},

	// no - ID can not contains more than 32 characters
	{
		[]byte("urn:abcdefghilmnopqrstuvzabcdefghilmn:specificstring"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 36]`,
		false,
	},

	// no - ID can not contain special characters
	{
		[]byte("URN:a!?:x"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 5]`,
		false,
	},
	{
		[]byte("URN:@,:x"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 4]`,
		false,
	},
	{
		[]byte("URN:#,:x"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 4]`,
		false,
	},
	{
		[]byte("URN:bc'.@:x"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 6]`,
		false,
	},

	// no - ID can not be equal to "urn"
	{
		[]byte("urn:urn:NSS"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to not contain the "urn" reserved string [col 7]`,
		false,
	},
	{
		[]byte("urn:URN:NSS"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to not contain the "urn" reserved string [col 7]`,
		false,
	},
	{
		[]byte("URN:URN:NSS"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to not contain the "urn" reserved string [col 7]`,
		false,
	},
	{
		[]byte("urn:UrN:NSS"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to not contain the "urn" reserved string [col 7]`,
		false,
	},
	{
		[]byte("urn:Urn:NSS"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to not contain the "urn" reserved string [col 7]`,
		false,
	},

	// no - ID can not contain spaces
	{
		[]byte("urn:white space:NSS"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 9]`,
		false,
	},

	// no - SS can not contain spaces
	{
		[]byte("urn:concat:no spaces"),
		false,
		nil,
		"",
		"",
		`expecting the specific string to be a string containing alnum, hex, or others ([()+,-.:=@;$_!*']) chars [col 13]`,
		false,
	},

	// no - SS can not contain reserved characters (can accept them only if %-escaped)
	{
		[]byte("urn:a:%"), // the presence of an "%" character in an URN MUST be followed by two characters from the <hex> character set
		false,
		nil,
		"",
		"",
		`expecting the specific string hex chars to be well-formed (%alnum{2}) [col 7]`,
		false,
	},
	{
		[]byte("urn:a:?"),
		false,
		nil,
		"",
		"",
		`expecting the specific string to be a string containing alnum, hex, or others ([()+,-.:=@;$_!*']) chars [col 6]`,
		false,
	},
	{
		[]byte("urn:a:#"),
		false,
		nil,
		"",
		"",
		`expecting the specific string to be a string containing alnum, hex, or others ([()+,-.:=@;$_!*']) chars [col 6]`,
		false,
	},
	{
		[]byte("urn:a:/"),
		false,
		nil,
		"",
		"",
		`expecting the specific string to be a string containing alnum, hex, or others ([()+,-.:=@;$_!*']) chars [col 6]`,
		false,
	},

	// no - Incomplete URNs
	{
		[]byte("urn:"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 4]`,
		false,
	},
	{
		[]byte("urn::"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 4]`,
		false,
	},
	{
		[]byte("urn:a:"),
		false,
		nil,
		"",
		"",
		`expecting the specific string to be a string containing alnum, hex, or others ([()+,-.:=@;$_!*']) chars [col 6]`,
		false,
	},
	// {
	// 	"urn:a",
	// 	false,
	// 	nil,
	// 	"",
	// 	"",
	// 	"",
	// },
}

var equivalenceTests = []struct {
	eq bool
	lx []byte
	rx []byte
}{
	{
		true,
		[]byte("urn:foo:a123%2C456"),
		[]byte("URN:FOO:a123%2c456"),
	},
	{
		true,
		[]byte("urn:foo:AbC123%2C456"),
		[]byte("URN:FOO:AbC123%2c456"),
	},
	{
		true,
		[]byte("urn:foo:AbC123%2C456%1f"),
		[]byte("URN:FOO:AbC123%2c456%1f"),
	},
	{
		true,
		[]byte("URN:foo:a123,456"),
		[]byte("urn:foo:a123,456"),
	},
	{
		true,
		[]byte("URN:foo:a123,456"),
		[]byte("urn:FOO:a123,456"),
	},
	{
		true,
		[]byte("urn:foo:a123,456"),
		[]byte("urn:FOO:a123,456"),
	},
	{
		true,
		[]byte("urn:ciao:%2E"),
		[]byte("urn:ciao:%2e"),
	},
	{
		false,
		[]byte("urn:foo:A123,456"),
		[]byte("URN:foo:a123,456"),
	},
	{
		false,
		[]byte("urn:foo:A123,456"),
		[]byte("urn:foo:a123,456"),
	},
	{
		false,
		[]byte("urn:foo:A123,456"),
		[]byte("urn:FOO:a123,456"),
	},
}

var fallbackTestCases = []testCase{
	// ok SCIM
	{
		[]byte("urn:ietf:params:scim:schemas:core:2.0:User"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "schemas:core:2.0:User",
		},
		"urn:ietf:params:scim:schemas:core:2.0:User",
		"urn:ietf:params:scim:schemas:core:2.0:User",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "schemas:extension:enterprise:2.0:User",
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "schemas:extension:enterprise:2.0:User:userName",
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName",
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:api:messages:3.0:Get"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "api:messages:3.0:Get",
		},
		"urn:ietf:params:scim:api:messages:3.0:Get",
		"urn:ietf:params:scim:api:messages:3.0:Get",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:core"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "schemas:core",
		},
		"urn:ietf:params:scim:schemas:core",
		"urn:ietf:params:scim:schemas:core",
		"",
		true,
	},
	{
		[]byte("urn:ietf:params:scim:param:core"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf:params:scim",
			SS:     "param:core",
		},
		"urn:ietf:params:scim:param:core",
		"urn:ietf:params:scim:param:core",
		"",
		true,
	},
	// no SCIM, ok URN
	{
		[]byte("urn:simple:ciao"),
		true,
		&URN{
			prefix: "urn",
			ID:     "simple",
			SS:     "ciao",
		},
		"urn:simple:ciao",
		"urn:simple:ciao",
		"",
		false,
	},
	{
		[]byte("urn:WRONG4SCIM:schemas:core"),
		true,
		&URN{
			prefix: "urn",
			ID:     "WRONG4SCIM",
			SS:     "schemas:core",
		},
		"urn:WRONG4SCIM:schemas:core",
		"urn:wrong4scim:schemas:core",
		"",
		false,
	},
	{
		[]byte("urn:ietf:params:scim:ERR:core"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf",
			SS:     "params:scim:ERR:core",
		},
		"urn:ietf:params:scim:ERR:core",
		"urn:ietf:params:scim:ERR:core",
		"",
		false,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:$"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf",
			SS:     "params:scim:schemas:$",
		},
		"urn:ietf:params:scim:schemas:$",
		"urn:ietf:params:scim:schemas:$",
		"",
		false,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:core-"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf",
			SS:     "params:scim:schemas:core-",
		},
		"urn:ietf:params:scim:schemas:core-",
		"urn:ietf:params:scim:schemas:core-",
		"",
		false,
	},
	{
		[]byte("urn:ietf:params:scim:api:core:"),
		true,
		&URN{
			prefix: "urn",
			ID:     "ietf",
			SS:     "params:scim:api:core:",
		},
		"urn:ietf:params:scim:api:core:",
		"urn:ietf:params:scim:api:core:",
		"",
		false,
	},
	// no SCIM, no URN
	{
		[]byte("arn:ietf:params:scim:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 0),
		false,
	},
	{
		[]byte("usn:ietf:params:scim:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 1),
		false,
	},
	{
		[]byte("urm:ietf:params:scim:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 2),
		false,
	},
	{
		[]byte("urno:ietf:params:scim:schemas:core"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 3),
		false,
	},
	{
		[]byte("urno"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errPrefix, 3),
		false,
	},
	{
		[]byte("URN:a!?:x"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 5]`,
		false,
	},
	{
		[]byte("urn:Urn:NSS"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to not contain the "urn" reserved string [col 7]`,
		false,
	},
	{
		[]byte("urn:spazio bianco:NSS"),
		false,
		nil,
		"",
		"",
		`expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col 10]`,
		false,
	},
	{
		[]byte("urn:conca:z ws"),
		false,
		nil,
		"",
		"",
		`expecting the specific string to be a string containing alnum, hex, or others ([()+,-.:=@;$_!*']) chars [col 11]`,
		false,
	},
	{
		[]byte("urn:ietf:params:scim:schemas:core:2.&"),
		false,
		nil,
		"",
		"",
		fmt.Sprintf(errSpecificString, 36),
		false,
	},
}
