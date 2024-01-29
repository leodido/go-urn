package urn

import (
	"fmt"

	scimschema "github.com/leodido/go-urn/scim/schema"
)

var (
	errPrefix              = "expecting the prefix to be the \"urn\" string (whatever case) [col %d]"
	errIdentifier          = "expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its beginning) [col %d]"
	errSpecificString      = "expecting the specific string to be a string containing alnum, hex, or others ([()+,-.:=@;$_!*']) chars [col %d]"
	errNoUrnWithinID       = "expecting the identifier to not contain the \"urn\" reserved string [col %d]"
	errHex                 = "expecting the specific string hex chars to be well-formed (%%alnum{2}) [col %d]"
	errParse               = "parsing error [col %d]"
	errSCIMNamespace       = "expecing the SCIM namespace identifier (ietf:params:scim) [col %d]"
	errSCIMType            = "expecting a correct SCIM type (schemas, api, param) [col %d]"
	errSCIMName            = "expecting one or more alnum char in the SCIM name part [col %d]"
	errSCIMOther           = "expecting a well-formed other SCIM part [col %d]"
	errSCIMOtherIncomplete = "expecting a not empty SCIM other part after colon [col %d]"
	err8141InformalID      = "informal URN namespace must be in the form urn-[1-9][0-9] [col %d]"
	err8141SpecificString  = "expecting the specific string to contain alnum, hex, or others ([~&()+,-.:=@;$_!*'] or [/?] not in first position) chars [col %d]"
	err8141Identifier      = "expecting the indentifier to be a string with (length 2 to 32 chars) containing alnum (or dashes) not starting or ending with a dash [col %d]"
	err8141RComponentStart = "expecting only one r-component (starting with the ?+ sequence) [col %d]"
	err8141QComponentStart = "expecting only one q-component (starting with the ?= sequence) [col %d]"
	err8141MalformedRComp  = "expecting a non-empty r-component containing alnum, hex, or others ([~&()+,-.:=@;$_!*'] or [/?] but not at its beginning) [col %d]"
	err8141MalformedQComp  = "expecting a non-empty q-component containing alnum, hex, or others ([~&()+,-.:=@;$_!*'] or [/?] but not at its beginning) [col %d]"
)

const start int = 1
const firstFinal int = 176

const enUrn int = 5
const enUrnOnly int = 44
const enScim int = 48
const enScimOnly int = 83
const enRfc8141 int = 87
const enRfc8141Only int = 172
const enFail int = 201
const enMain int = 1

// Machine is the interface representing the FSM
type Machine interface {
	Error() error
	Parse(input []byte) (*URN, error)
	WithParsingMode(ParsingMode)
}

type machine struct {
	data           []byte
	cs             int
	p, pe, eof, pb int
	err            error
	tolower        []int
	parsingMode    ParsingMode
	parsingModeSet bool
}

// NewMachine creates a new FSM able to parse RFC 2141 strings.
func NewMachine(options ...Option) Machine {
	m := &machine{
		parsingModeSet: false,
	}

	for _, o := range options {
		o(m)
	}
	// Set default parsing mode
	if !m.parsingModeSet {
		m.WithParsingMode(DefaultParsingMode)
	}

	return m
}

// Err returns the error that occurred on the last call to Parse.
//
// If the result is nil, then the line was parsed successfully.
func (m *machine) Error() error {
	return m.err
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

// Parse parses the input byte array as a RFC 2141 or RFC7643 string.
func (m *machine) Parse(input []byte) (*URN, error) {
	m.data = input
	m.p = 0
	m.pb = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil
	m.tolower = []int{}
	output := &URN{}

	switch m.parsingMode {
	case RFC2141Only:
		m.cs = enUrnOnly

	case RFC7643Only:
		m.cs = enScimOnly

	case RFC8141Only:
		m.cs = enRfc8141Only

	case All:
		fallthrough
	default:
		{
			m.cs = start
		}
	}
	{
		if (m.p) == (m.pe) {
			goto _testEof
		}
		switch m.cs {
		case 1:
			goto stCase1
		case 0:
			goto stCase0
		case 2:
			goto stCase2
		case 3:
			goto stCase3
		case 4:
			goto stCase4
		case 176:
			goto stCase176
		case 5:
			goto stCase5
		case 6:
			goto stCase6
		case 7:
			goto stCase7
		case 8:
			goto stCase8
		case 9:
			goto stCase9
		case 10:
			goto stCase10
		case 11:
			goto stCase11
		case 12:
			goto stCase12
		case 13:
			goto stCase13
		case 14:
			goto stCase14
		case 15:
			goto stCase15
		case 16:
			goto stCase16
		case 17:
			goto stCase17
		case 18:
			goto stCase18
		case 19:
			goto stCase19
		case 20:
			goto stCase20
		case 21:
			goto stCase21
		case 22:
			goto stCase22
		case 23:
			goto stCase23
		case 24:
			goto stCase24
		case 25:
			goto stCase25
		case 26:
			goto stCase26
		case 27:
			goto stCase27
		case 28:
			goto stCase28
		case 29:
			goto stCase29
		case 30:
			goto stCase30
		case 31:
			goto stCase31
		case 32:
			goto stCase32
		case 33:
			goto stCase33
		case 34:
			goto stCase34
		case 35:
			goto stCase35
		case 36:
			goto stCase36
		case 37:
			goto stCase37
		case 38:
			goto stCase38
		case 177:
			goto stCase177
		case 39:
			goto stCase39
		case 40:
			goto stCase40
		case 178:
			goto stCase178
		case 41:
			goto stCase41
		case 42:
			goto stCase42
		case 43:
			goto stCase43
		case 44:
			goto stCase44
		case 45:
			goto stCase45
		case 46:
			goto stCase46
		case 47:
			goto stCase47
		case 179:
			goto stCase179
		case 48:
			goto stCase48
		case 49:
			goto stCase49
		case 50:
			goto stCase50
		case 51:
			goto stCase51
		case 52:
			goto stCase52
		case 53:
			goto stCase53
		case 54:
			goto stCase54
		case 55:
			goto stCase55
		case 56:
			goto stCase56
		case 57:
			goto stCase57
		case 58:
			goto stCase58
		case 59:
			goto stCase59
		case 60:
			goto stCase60
		case 61:
			goto stCase61
		case 62:
			goto stCase62
		case 63:
			goto stCase63
		case 64:
			goto stCase64
		case 65:
			goto stCase65
		case 66:
			goto stCase66
		case 67:
			goto stCase67
		case 68:
			goto stCase68
		case 69:
			goto stCase69
		case 180:
			goto stCase180
		case 70:
			goto stCase70
		case 181:
			goto stCase181
		case 71:
			goto stCase71
		case 72:
			goto stCase72
		case 182:
			goto stCase182
		case 73:
			goto stCase73
		case 74:
			goto stCase74
		case 75:
			goto stCase75
		case 76:
			goto stCase76
		case 77:
			goto stCase77
		case 78:
			goto stCase78
		case 79:
			goto stCase79
		case 80:
			goto stCase80
		case 81:
			goto stCase81
		case 82:
			goto stCase82
		case 83:
			goto stCase83
		case 84:
			goto stCase84
		case 85:
			goto stCase85
		case 86:
			goto stCase86
		case 183:
			goto stCase183
		case 87:
			goto stCase87
		case 88:
			goto stCase88
		case 89:
			goto stCase89
		case 90:
			goto stCase90
		case 91:
			goto stCase91
		case 92:
			goto stCase92
		case 93:
			goto stCase93
		case 94:
			goto stCase94
		case 95:
			goto stCase95
		case 96:
			goto stCase96
		case 97:
			goto stCase97
		case 98:
			goto stCase98
		case 99:
			goto stCase99
		case 100:
			goto stCase100
		case 101:
			goto stCase101
		case 102:
			goto stCase102
		case 103:
			goto stCase103
		case 104:
			goto stCase104
		case 105:
			goto stCase105
		case 106:
			goto stCase106
		case 107:
			goto stCase107
		case 108:
			goto stCase108
		case 109:
			goto stCase109
		case 110:
			goto stCase110
		case 111:
			goto stCase111
		case 112:
			goto stCase112
		case 113:
			goto stCase113
		case 114:
			goto stCase114
		case 115:
			goto stCase115
		case 116:
			goto stCase116
		case 117:
			goto stCase117
		case 118:
			goto stCase118
		case 119:
			goto stCase119
		case 120:
			goto stCase120
		case 184:
			goto stCase184
		case 185:
			goto stCase185
		case 186:
			goto stCase186
		case 121:
			goto stCase121
		case 122:
			goto stCase122
		case 187:
			goto stCase187
		case 123:
			goto stCase123
		case 124:
			goto stCase124
		case 188:
			goto stCase188
		case 125:
			goto stCase125
		case 126:
			goto stCase126
		case 189:
			goto stCase189
		case 127:
			goto stCase127
		case 128:
			goto stCase128
		case 190:
			goto stCase190
		case 191:
			goto stCase191
		case 192:
			goto stCase192
		case 193:
			goto stCase193
		case 194:
			goto stCase194
		case 129:
			goto stCase129
		case 130:
			goto stCase130
		case 195:
			goto stCase195
		case 196:
			goto stCase196
		case 197:
			goto stCase197
		case 131:
			goto stCase131
		case 132:
			goto stCase132
		case 133:
			goto stCase133
		case 198:
			goto stCase198
		case 134:
			goto stCase134
		case 135:
			goto stCase135
		case 136:
			goto stCase136
		case 199:
			goto stCase199
		case 137:
			goto stCase137
		case 138:
			goto stCase138
		case 139:
			goto stCase139
		case 140:
			goto stCase140
		case 141:
			goto stCase141
		case 142:
			goto stCase142
		case 143:
			goto stCase143
		case 144:
			goto stCase144
		case 145:
			goto stCase145
		case 146:
			goto stCase146
		case 147:
			goto stCase147
		case 148:
			goto stCase148
		case 149:
			goto stCase149
		case 150:
			goto stCase150
		case 151:
			goto stCase151
		case 152:
			goto stCase152
		case 153:
			goto stCase153
		case 154:
			goto stCase154
		case 155:
			goto stCase155
		case 156:
			goto stCase156
		case 157:
			goto stCase157
		case 158:
			goto stCase158
		case 159:
			goto stCase159
		case 160:
			goto stCase160
		case 161:
			goto stCase161
		case 162:
			goto stCase162
		case 163:
			goto stCase163
		case 164:
			goto stCase164
		case 165:
			goto stCase165
		case 166:
			goto stCase166
		case 167:
			goto stCase167
		case 168:
			goto stCase168
		case 169:
			goto stCase169
		case 170:
			goto stCase170
		case 171:
			goto stCase171
		case 172:
			goto stCase172
		case 173:
			goto stCase173
		case 174:
			goto stCase174
		case 175:
			goto stCase175
		case 200:
			goto stCase200
		case 201:
			goto stCase201
		}
		goto stOut
	stCase1:
		switch (m.data)[(m.p)] {
		case 85:
			goto tr1
		case 117:
			goto tr1
		}
		goto tr0
	tr0:

		m.err = fmt.Errorf(errPrefix, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errParse, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr5:

		m.err = fmt.Errorf(errIdentifier, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errPrefix, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errParse, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr8:

		m.err = fmt.Errorf(errIdentifier, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errParse, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr41:

		m.err = fmt.Errorf(errSpecificString, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errParse, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr44:

		if m.parsingMode == RFC2141Only || m.parsingMode == All {
			m.err = fmt.Errorf(errHex, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		// Otherwise, we expect the machine to fallback to SCIM errors

		m.err = fmt.Errorf(errSpecificString, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errParse, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr51:

		m.err = fmt.Errorf(errIdentifier, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errNoUrnWithinID, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errParse, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr52:

		m.err = fmt.Errorf(errPrefix, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr57:

		// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
		if m.parsingMode == All {
			// TODO: store why the machine fallback to the RFC2141 one?
			output.scim = nil
			// Rewind the cursor after the prefix ends ("urn:")
			(m.p) = (4) - 1

			// Go to the "urn" machine from this point on
			{
				goto st5
			}
		}
		m.err = fmt.Errorf(errSCIMNamespace, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr75:

		// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
		if m.parsingMode == All {
			// TODO: store why the machine fallback to the RFC2141 one?
			output.scim = nil
			// Rewind the cursor after the prefix ends ("urn:")
			(m.p) = (4) - 1

			// Go to the "urn" machine from this point on
			{
				goto st5
			}
		}
		m.err = fmt.Errorf(errSCIMType, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr82:

		// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
		if m.parsingMode == All {
			// TODO: store why the machine fallback to the RFC2141 one?
			output.scim = nil
			// Rewind the cursor after the prefix ends ("urn:")
			(m.p) = (4) - 1

			// Go to the "urn" machine from this point on
			{
				goto st5
			}
		}
		m.err = fmt.Errorf(errSCIMName, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr84:

		// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
		if m.parsingMode == All {
			// TODO: store why the machine fallback to the RFC2141 one?
			output.scim = nil
			// Rewind the cursor after the prefix ends ("urn:")
			(m.p) = (4) - 1

			// Go to the "urn" machine from this point on
			{
				goto st5
			}
		}
		if m.p == m.pe {
			m.err = fmt.Errorf(errSCIMOtherIncomplete, m.p-1)
		} else {
			m.err = fmt.Errorf(errSCIMOther, m.p)
		}
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr87:

		if m.parsingMode == RFC2141Only || m.parsingMode == All {
			m.err = fmt.Errorf(errHex, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		// Otherwise, we expect the machine to fallback to SCIM errors

		// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
		if m.parsingMode == All {
			// TODO: store why the machine fallback to the RFC2141 one?
			output.scim = nil
			// Rewind the cursor after the prefix ends ("urn:")
			(m.p) = (4) - 1

			// Go to the "urn" machine from this point on
			{
				goto st5
			}
		}
		if m.p == m.pe {
			m.err = fmt.Errorf(errSCIMOtherIncomplete, m.p-1)
		} else {
			m.err = fmt.Errorf(errSCIMOther, m.p)
		}
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr104:

		m.err = fmt.Errorf(err8141Identifier, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(errPrefix, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr107:

		m.err = fmt.Errorf(err8141Identifier, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr170:

		m.err = fmt.Errorf(err8141SpecificString, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr173:

		if m.parsingMode == RFC2141Only || m.parsingMode == All {
			m.err = fmt.Errorf(errHex, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		// Otherwise, we expect the machine to fallback to SCIM errors

		goto st0
	tr178:

		if m.parsingMode == RFC2141Only || m.parsingMode == All {
			m.err = fmt.Errorf(errHex, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		// Otherwise, we expect the machine to fallback to SCIM errors

		m.err = fmt.Errorf(err8141SpecificString, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr186:

		m.err = fmt.Errorf(err8141MalformedRComp, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr190:

		if m.parsingMode == RFC2141Only || m.parsingMode == All {
			m.err = fmt.Errorf(errHex, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		// Otherwise, we expect the machine to fallback to SCIM errors

		m.err = fmt.Errorf(err8141MalformedRComp, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr195:

		if m.parsingMode == RFC2141Only || m.parsingMode == All {
			m.err = fmt.Errorf(errHex, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		// Otherwise, we expect the machine to fallback to SCIM errors

		m.err = fmt.Errorf(err8141MalformedQComp, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr200:

		if m.parsingMode == RFC2141Only || m.parsingMode == All {
			m.err = fmt.Errorf(errHex, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		// Otherwise, we expect the machine to fallback to SCIM errors

		m.err = fmt.Errorf(err8141MalformedRComp, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(err8141MalformedQComp, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr204:

		m.err = fmt.Errorf(err8141MalformedQComp, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	tr241:

		m.err = fmt.Errorf(err8141MalformedRComp, m.p)
		(m.p)--

		{
			goto st201
		}

		m.err = fmt.Errorf(err8141MalformedQComp, m.p)
		(m.p)--

		{
			goto st201
		}

		goto st0
	stCase0:
	st0:
		m.cs = 0
		goto _out
	tr1:

		m.pb = m.p

		if m.parsingMode != RFC8141Only {
			// Throw an error when:
			// - we are entering here matching the the prefix in the namespace identifier part
			// - looking ahead (3 chars) we find a colon
			if pos := m.p + 3; pos < m.pe && m.data[pos] == 58 && output.prefix != "" {
				m.err = fmt.Errorf(errNoUrnWithinID, pos)
				(m.p)--

				{
					goto st201
				}
			}
		}

		goto st2
	st2:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof2
		}
	stCase2:
		switch (m.data)[(m.p)] {
		case 82:
			goto st3
		case 114:
			goto st3
		}
		goto tr0
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof3
		}
	stCase3:
		switch (m.data)[(m.p)] {
		case 78:
			goto st4
		case 110:
			goto st4
		}
		goto tr0
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof4
		}
	stCase4:
		if (m.data)[(m.p)] == 58 {
			goto tr4
		}
		goto tr0
	tr4:

		output.prefix = string(m.text())
		{
			goto st48
		}
		goto st176
	st176:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof176
		}
	stCase176:
		goto tr0
	st5:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof5
		}
	stCase5:
		switch (m.data)[(m.p)] {
		case 85:
			goto tr7
		case 117:
			goto tr7
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr6
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr6
			}
		default:
			goto tr6
		}
		goto tr5
	tr6:

		m.pb = m.p

		goto st6
	st6:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof6
		}
	stCase6:
		switch (m.data)[(m.p)] {
		case 45:
			goto st7
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st7
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st7
			}
		default:
			goto st7
		}
		goto tr8
	st7:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof7
		}
	stCase7:
		switch (m.data)[(m.p)] {
		case 45:
			goto st8
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st8
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr8
	st8:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof8
		}
	stCase8:
		switch (m.data)[(m.p)] {
		case 45:
			goto st9
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st9
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr8
	st9:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof9
		}
	stCase9:
		switch (m.data)[(m.p)] {
		case 45:
			goto st10
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st10
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st10
			}
		default:
			goto st10
		}
		goto tr8
	st10:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof10
		}
	stCase10:
		switch (m.data)[(m.p)] {
		case 45:
			goto st11
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st11
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr8
	st11:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof11
		}
	stCase11:
		switch (m.data)[(m.p)] {
		case 45:
			goto st12
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st12
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st12
			}
		default:
			goto st12
		}
		goto tr8
	st12:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof12
		}
	stCase12:
		switch (m.data)[(m.p)] {
		case 45:
			goto st13
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st13
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr8
	st13:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof13
		}
	stCase13:
		switch (m.data)[(m.p)] {
		case 45:
			goto st14
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st14
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr8
	st14:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof14
		}
	stCase14:
		switch (m.data)[(m.p)] {
		case 45:
			goto st15
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st15
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr8
	st15:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof15
		}
	stCase15:
		switch (m.data)[(m.p)] {
		case 45:
			goto st16
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st16
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr8
	st16:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof16
		}
	stCase16:
		switch (m.data)[(m.p)] {
		case 45:
			goto st17
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st17
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st17
			}
		default:
			goto st17
		}
		goto tr8
	st17:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof17
		}
	stCase17:
		switch (m.data)[(m.p)] {
		case 45:
			goto st18
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st18
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st18
			}
		default:
			goto st18
		}
		goto tr8
	st18:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof18
		}
	stCase18:
		switch (m.data)[(m.p)] {
		case 45:
			goto st19
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st19
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st19
			}
		default:
			goto st19
		}
		goto tr8
	st19:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof19
		}
	stCase19:
		switch (m.data)[(m.p)] {
		case 45:
			goto st20
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st20
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st20
			}
		default:
			goto st20
		}
		goto tr8
	st20:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof20
		}
	stCase20:
		switch (m.data)[(m.p)] {
		case 45:
			goto st21
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st21
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st21
			}
		default:
			goto st21
		}
		goto tr8
	st21:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof21
		}
	stCase21:
		switch (m.data)[(m.p)] {
		case 45:
			goto st22
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st22
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st22
			}
		default:
			goto st22
		}
		goto tr8
	st22:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof22
		}
	stCase22:
		switch (m.data)[(m.p)] {
		case 45:
			goto st23
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st23
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto tr8
	st23:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof23
		}
	stCase23:
		switch (m.data)[(m.p)] {
		case 45:
			goto st24
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st24
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st24
			}
		default:
			goto st24
		}
		goto tr8
	st24:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof24
		}
	stCase24:
		switch (m.data)[(m.p)] {
		case 45:
			goto st25
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st25
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st25
			}
		default:
			goto st25
		}
		goto tr8
	st25:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof25
		}
	stCase25:
		switch (m.data)[(m.p)] {
		case 45:
			goto st26
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st26
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr8
	st26:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof26
		}
	stCase26:
		switch (m.data)[(m.p)] {
		case 45:
			goto st27
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st27
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st27
			}
		default:
			goto st27
		}
		goto tr8
	st27:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof27
		}
	stCase27:
		switch (m.data)[(m.p)] {
		case 45:
			goto st28
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st28
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st28
			}
		default:
			goto st28
		}
		goto tr8
	st28:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof28
		}
	stCase28:
		switch (m.data)[(m.p)] {
		case 45:
			goto st29
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st29
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st29
			}
		default:
			goto st29
		}
		goto tr8
	st29:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof29
		}
	stCase29:
		switch (m.data)[(m.p)] {
		case 45:
			goto st30
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st30
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st30
			}
		default:
			goto st30
		}
		goto tr8
	st30:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof30
		}
	stCase30:
		switch (m.data)[(m.p)] {
		case 45:
			goto st31
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st31
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st31
			}
		default:
			goto st31
		}
		goto tr8
	st31:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof31
		}
	stCase31:
		switch (m.data)[(m.p)] {
		case 45:
			goto st32
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st32
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st32
			}
		default:
			goto st32
		}
		goto tr8
	st32:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof32
		}
	stCase32:
		switch (m.data)[(m.p)] {
		case 45:
			goto st33
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st33
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st33
			}
		default:
			goto st33
		}
		goto tr8
	st33:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof33
		}
	stCase33:
		switch (m.data)[(m.p)] {
		case 45:
			goto st34
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st34
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st34
			}
		default:
			goto st34
		}
		goto tr8
	st34:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof34
		}
	stCase34:
		switch (m.data)[(m.p)] {
		case 45:
			goto st35
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st35
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st35
			}
		default:
			goto st35
		}
		goto tr8
	st35:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof35
		}
	stCase35:
		switch (m.data)[(m.p)] {
		case 45:
			goto st36
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st36
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st36
			}
		default:
			goto st36
		}
		goto tr8
	st36:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof36
		}
	stCase36:
		switch (m.data)[(m.p)] {
		case 45:
			goto st37
		case 58:
			goto tr10
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st37
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto tr8
	st37:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof37
		}
	stCase37:
		if (m.data)[(m.p)] == 58 {
			goto tr10
		}
		goto tr8
	tr10:

		output.ID = string(m.text())

		goto st38
	st38:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof38
		}
	stCase38:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr42
		case 36:
			goto tr42
		case 37:
			goto tr43
		case 61:
			goto tr42
		case 95:
			goto tr42
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 39 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto tr42
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr42
				}
			case (m.data)[(m.p)] >= 64:
				goto tr42
			}
		default:
			goto tr42
		}
		goto tr41
	tr42:

		m.pb = m.p

		goto st177
	st177:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof177
		}
	stCase177:
		switch (m.data)[(m.p)] {
		case 33:
			goto st177
		case 36:
			goto st177
		case 37:
			goto st39
		case 61:
			goto st177
		case 95:
			goto st177
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 39 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto st177
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st177
				}
			case (m.data)[(m.p)] >= 64:
				goto st177
			}
		default:
			goto st177
		}
		goto tr41
	tr43:

		m.pb = m.p

		goto st39
	st39:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof39
		}
	stCase39:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st40
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st40
			}
		default:
			goto tr46
		}
		goto tr44
	tr46:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st40
	st40:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof40
		}
	stCase40:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st178
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st178
			}
		default:
			goto tr48
		}
		goto tr44
	tr48:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st178
	st178:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof178
		}
	stCase178:
		switch (m.data)[(m.p)] {
		case 33:
			goto st177
		case 36:
			goto st177
		case 37:
			goto st39
		case 61:
			goto st177
		case 95:
			goto st177
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 39 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto st177
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st177
				}
			case (m.data)[(m.p)] >= 64:
				goto st177
			}
		default:
			goto st177
		}
		goto tr44
	tr7:

		m.pb = m.p

		if m.parsingMode != RFC8141Only {
			// Throw an error when:
			// - we are entering here matching the the prefix in the namespace identifier part
			// - looking ahead (3 chars) we find a colon
			if pos := m.p + 3; pos < m.pe && m.data[pos] == 58 && output.prefix != "" {
				m.err = fmt.Errorf(errNoUrnWithinID, pos)
				(m.p)--

				{
					goto st201
				}
			}
		}

		goto st41
	st41:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof41
		}
	stCase41:
		switch (m.data)[(m.p)] {
		case 45:
			goto st7
		case 58:
			goto tr10
		case 82:
			goto st42
		case 114:
			goto st42
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st7
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st7
			}
		default:
			goto st7
		}
		goto tr5
	st42:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof42
		}
	stCase42:
		switch (m.data)[(m.p)] {
		case 45:
			goto st8
		case 58:
			goto tr10
		case 78:
			goto st43
		case 110:
			goto st43
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st8
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr5
	st43:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof43
		}
	stCase43:
		if (m.data)[(m.p)] == 45 {
			goto st9
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st9
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr51
	stCase44:
		switch (m.data)[(m.p)] {
		case 85:
			goto tr53
		case 117:
			goto tr53
		}
		goto tr52
	tr53:

		m.pb = m.p

		if m.parsingMode != RFC8141Only {
			// Throw an error when:
			// - we are entering here matching the the prefix in the namespace identifier part
			// - looking ahead (3 chars) we find a colon
			if pos := m.p + 3; pos < m.pe && m.data[pos] == 58 && output.prefix != "" {
				m.err = fmt.Errorf(errNoUrnWithinID, pos)
				(m.p)--

				{
					goto st201
				}
			}
		}

		goto st45
	st45:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof45
		}
	stCase45:
		switch (m.data)[(m.p)] {
		case 82:
			goto st46
		case 114:
			goto st46
		}
		goto tr52
	st46:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof46
		}
	stCase46:
		switch (m.data)[(m.p)] {
		case 78:
			goto st47
		case 110:
			goto st47
		}
		goto tr52
	st47:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof47
		}
	stCase47:
		if (m.data)[(m.p)] == 58 {
			goto tr56
		}
		goto tr52
	tr56:

		output.prefix = string(m.text())
		{
			goto st5
		}
		goto st179
	st179:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof179
		}
	stCase179:
		goto tr52
	st48:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof48
		}
	stCase48:
		if (m.data)[(m.p)] == 105 {
			goto tr58
		}
		goto tr57
	tr58:

		m.pb = m.p

		goto st49
	st49:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof49
		}
	stCase49:
		if (m.data)[(m.p)] == 101 {
			goto st50
		}
		goto tr57
	st50:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof50
		}
	stCase50:
		if (m.data)[(m.p)] == 116 {
			goto st51
		}
		goto tr57
	st51:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof51
		}
	stCase51:
		if (m.data)[(m.p)] == 102 {
			goto st52
		}
		goto tr57
	st52:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof52
		}
	stCase52:
		if (m.data)[(m.p)] == 58 {
			goto st53
		}
		goto tr57
	st53:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof53
		}
	stCase53:
		if (m.data)[(m.p)] == 112 {
			goto st54
		}
		goto tr57
	st54:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof54
		}
	stCase54:
		if (m.data)[(m.p)] == 97 {
			goto st55
		}
		goto tr57
	st55:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof55
		}
	stCase55:
		if (m.data)[(m.p)] == 114 {
			goto st56
		}
		goto tr57
	st56:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof56
		}
	stCase56:
		if (m.data)[(m.p)] == 97 {
			goto st57
		}
		goto tr57
	st57:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof57
		}
	stCase57:
		if (m.data)[(m.p)] == 109 {
			goto st58
		}
		goto tr57
	st58:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof58
		}
	stCase58:
		if (m.data)[(m.p)] == 115 {
			goto st59
		}
		goto tr57
	st59:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof59
		}
	stCase59:
		if (m.data)[(m.p)] == 58 {
			goto st60
		}
		goto tr57
	st60:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof60
		}
	stCase60:
		if (m.data)[(m.p)] == 115 {
			goto st61
		}
		goto tr57
	st61:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof61
		}
	stCase61:
		if (m.data)[(m.p)] == 99 {
			goto st62
		}
		goto tr57
	st62:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof62
		}
	stCase62:
		if (m.data)[(m.p)] == 105 {
			goto st63
		}
		goto tr57
	st63:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof63
		}
	stCase63:
		if (m.data)[(m.p)] == 109 {
			goto st64
		}
		goto tr57
	st64:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof64
		}
	stCase64:
		if (m.data)[(m.p)] == 58 {
			goto tr74
		}
		goto tr57
	tr74:

		output.ID = string(m.text())

		output.scim = &SCIM{}

		goto st65
	st65:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof65
		}
	stCase65:
		switch (m.data)[(m.p)] {
		case 97:
			goto tr76
		case 112:
			goto tr77
		case 115:
			goto tr78
		}
		goto tr75
	tr76:

		m.pb = m.p

		goto st66
	st66:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof66
		}
	stCase66:
		if (m.data)[(m.p)] == 112 {
			goto st67
		}
		goto tr75
	st67:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof67
		}
	stCase67:
		if (m.data)[(m.p)] == 105 {
			goto st68
		}
		goto tr75
	st68:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof68
		}
	stCase68:
		if (m.data)[(m.p)] == 58 {
			goto tr81
		}
		goto tr75
	tr81:

		output.scim.Type = scimschema.TypeFromString(string(m.text()))

		goto st69
	st69:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof69
		}
	stCase69:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr83
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr83
			}
		default:
			goto tr83
		}
		goto tr82
	tr83:

		output.scim.pos = m.p

		goto st180
	st180:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof180
		}
	stCase180:
		if (m.data)[(m.p)] == 58 {
			goto tr220
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st180
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st180
			}
		default:
			goto st180
		}
		goto tr82
	tr220:

		output.scim.Name = string(m.data[output.scim.pos:m.p])

		goto st70
	st70:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof70
		}
	stCase70:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr85
		case 36:
			goto tr85
		case 37:
			goto tr86
		case 61:
			goto tr85
		case 95:
			goto tr85
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 39 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto tr85
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr85
				}
			case (m.data)[(m.p)] >= 64:
				goto tr85
			}
		default:
			goto tr85
		}
		goto tr84
	tr85:

		output.scim.pos = m.p

		goto st181
	st181:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof181
		}
	stCase181:
		switch (m.data)[(m.p)] {
		case 33:
			goto st181
		case 36:
			goto st181
		case 37:
			goto st71
		case 61:
			goto st181
		case 95:
			goto st181
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 39 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto st181
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st181
				}
			case (m.data)[(m.p)] >= 64:
				goto st181
			}
		default:
			goto st181
		}
		goto tr84
	tr86:

		output.scim.pos = m.p

		goto st71
	st71:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof71
		}
	stCase71:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st72
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st72
			}
		default:
			goto tr89
		}
		goto tr87
	tr89:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st72
	st72:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof72
		}
	stCase72:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st182
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st182
			}
		default:
			goto tr91
		}
		goto tr87
	tr91:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st182
	st182:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof182
		}
	stCase182:
		switch (m.data)[(m.p)] {
		case 33:
			goto st181
		case 36:
			goto st181
		case 37:
			goto st71
		case 61:
			goto st181
		case 95:
			goto st181
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 39 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto st181
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st181
				}
			case (m.data)[(m.p)] >= 64:
				goto st181
			}
		default:
			goto st181
		}
		goto tr87
	tr77:

		m.pb = m.p

		goto st73
	st73:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof73
		}
	stCase73:
		if (m.data)[(m.p)] == 97 {
			goto st74
		}
		goto tr75
	st74:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof74
		}
	stCase74:
		if (m.data)[(m.p)] == 114 {
			goto st75
		}
		goto tr75
	st75:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof75
		}
	stCase75:
		if (m.data)[(m.p)] == 97 {
			goto st76
		}
		goto tr75
	st76:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof76
		}
	stCase76:
		if (m.data)[(m.p)] == 109 {
			goto st68
		}
		goto tr75
	tr78:

		m.pb = m.p

		goto st77
	st77:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof77
		}
	stCase77:
		if (m.data)[(m.p)] == 99 {
			goto st78
		}
		goto tr75
	st78:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof78
		}
	stCase78:
		if (m.data)[(m.p)] == 104 {
			goto st79
		}
		goto tr75
	st79:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof79
		}
	stCase79:
		if (m.data)[(m.p)] == 101 {
			goto st80
		}
		goto tr75
	st80:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof80
		}
	stCase80:
		if (m.data)[(m.p)] == 109 {
			goto st81
		}
		goto tr75
	st81:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof81
		}
	stCase81:
		if (m.data)[(m.p)] == 97 {
			goto st82
		}
		goto tr75
	st82:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof82
		}
	stCase82:
		if (m.data)[(m.p)] == 115 {
			goto st68
		}
		goto tr75
	stCase83:
		switch (m.data)[(m.p)] {
		case 85:
			goto tr100
		case 117:
			goto tr100
		}
		goto tr52
	tr100:

		m.pb = m.p

		if m.parsingMode != RFC8141Only {
			// Throw an error when:
			// - we are entering here matching the the prefix in the namespace identifier part
			// - looking ahead (3 chars) we find a colon
			if pos := m.p + 3; pos < m.pe && m.data[pos] == 58 && output.prefix != "" {
				m.err = fmt.Errorf(errNoUrnWithinID, pos)
				(m.p)--

				{
					goto st201
				}
			}
		}

		goto st84
	st84:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof84
		}
	stCase84:
		switch (m.data)[(m.p)] {
		case 82:
			goto st85
		case 114:
			goto st85
		}
		goto tr52
	st85:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof85
		}
	stCase85:
		switch (m.data)[(m.p)] {
		case 78:
			goto st86
		case 110:
			goto st86
		}
		goto tr52
	st86:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof86
		}
	stCase86:
		if (m.data)[(m.p)] == 58 {
			goto tr103
		}
		goto tr52
	tr103:

		output.prefix = string(m.text())
		{
			goto st48
		}
		goto st183
	st183:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof183
		}
	stCase183:
		goto tr52
	st87:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof87
		}
	stCase87:
		switch (m.data)[(m.p)] {
		case 85:
			goto tr106
		case 117:
			goto tr106
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr105
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr105
			}
		default:
			goto tr105
		}
		goto tr104
	tr105:

		m.pb = m.p

		goto st88
	st88:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof88
		}
	stCase88:
		if (m.data)[(m.p)] == 45 {
			goto st89
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st166
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st166
			}
		default:
			goto st166
		}
		goto tr107
	st89:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof89
		}
	stCase89:
		if (m.data)[(m.p)] == 45 {
			goto st90
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st165
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st165
			}
		default:
			goto st165
		}
		goto tr107
	st90:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof90
		}
	stCase90:
		if (m.data)[(m.p)] == 45 {
			goto st91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st164
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st164
			}
		default:
			goto st164
		}
		goto tr107
	st91:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof91
		}
	stCase91:
		if (m.data)[(m.p)] == 45 {
			goto st92
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st163
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st163
			}
		default:
			goto st163
		}
		goto tr107
	st92:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof92
		}
	stCase92:
		if (m.data)[(m.p)] == 45 {
			goto st93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st162
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st162
			}
		default:
			goto st162
		}
		goto tr107
	st93:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof93
		}
	stCase93:
		if (m.data)[(m.p)] == 45 {
			goto st94
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st161
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st161
			}
		default:
			goto st161
		}
		goto tr107
	st94:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof94
		}
	stCase94:
		if (m.data)[(m.p)] == 45 {
			goto st95
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st160
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st160
			}
		default:
			goto st160
		}
		goto tr107
	st95:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof95
		}
	stCase95:
		if (m.data)[(m.p)] == 45 {
			goto st96
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st159
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st159
			}
		default:
			goto st159
		}
		goto tr107
	st96:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof96
		}
	stCase96:
		if (m.data)[(m.p)] == 45 {
			goto st97
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st158
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st158
			}
		default:
			goto st158
		}
		goto tr107
	st97:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof97
		}
	stCase97:
		if (m.data)[(m.p)] == 45 {
			goto st98
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st157
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st157
			}
		default:
			goto st157
		}
		goto tr107
	st98:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof98
		}
	stCase98:
		if (m.data)[(m.p)] == 45 {
			goto st99
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st156
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st156
			}
		default:
			goto st156
		}
		goto tr107
	st99:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof99
		}
	stCase99:
		if (m.data)[(m.p)] == 45 {
			goto st100
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st155
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st155
			}
		default:
			goto st155
		}
		goto tr107
	st100:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof100
		}
	stCase100:
		if (m.data)[(m.p)] == 45 {
			goto st101
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st154
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st154
			}
		default:
			goto st154
		}
		goto tr107
	st101:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof101
		}
	stCase101:
		if (m.data)[(m.p)] == 45 {
			goto st102
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st153
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st153
			}
		default:
			goto st153
		}
		goto tr107
	st102:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof102
		}
	stCase102:
		if (m.data)[(m.p)] == 45 {
			goto st103
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st152
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st152
			}
		default:
			goto st152
		}
		goto tr107
	st103:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof103
		}
	stCase103:
		if (m.data)[(m.p)] == 45 {
			goto st104
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st151
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st151
			}
		default:
			goto st151
		}
		goto tr107
	st104:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof104
		}
	stCase104:
		if (m.data)[(m.p)] == 45 {
			goto st105
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st150
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st150
			}
		default:
			goto st150
		}
		goto tr107
	st105:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof105
		}
	stCase105:
		if (m.data)[(m.p)] == 45 {
			goto st106
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st149
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st149
			}
		default:
			goto st149
		}
		goto tr107
	st106:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof106
		}
	stCase106:
		if (m.data)[(m.p)] == 45 {
			goto st107
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st148
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto tr107
	st107:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof107
		}
	stCase107:
		if (m.data)[(m.p)] == 45 {
			goto st108
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st147
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st147
			}
		default:
			goto st147
		}
		goto tr107
	st108:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof108
		}
	stCase108:
		if (m.data)[(m.p)] == 45 {
			goto st109
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st146
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st146
			}
		default:
			goto st146
		}
		goto tr107
	st109:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof109
		}
	stCase109:
		if (m.data)[(m.p)] == 45 {
			goto st110
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st145
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st145
			}
		default:
			goto st145
		}
		goto tr107
	st110:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof110
		}
	stCase110:
		if (m.data)[(m.p)] == 45 {
			goto st111
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st144
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st144
			}
		default:
			goto st144
		}
		goto tr107
	st111:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof111
		}
	stCase111:
		if (m.data)[(m.p)] == 45 {
			goto st112
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st143
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st143
			}
		default:
			goto st143
		}
		goto tr107
	st112:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof112
		}
	stCase112:
		if (m.data)[(m.p)] == 45 {
			goto st113
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st142
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st142
			}
		default:
			goto st142
		}
		goto tr107
	st113:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof113
		}
	stCase113:
		if (m.data)[(m.p)] == 45 {
			goto st114
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st141
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st141
			}
		default:
			goto st141
		}
		goto tr107
	st114:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof114
		}
	stCase114:
		if (m.data)[(m.p)] == 45 {
			goto st115
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st140
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st140
			}
		default:
			goto st140
		}
		goto tr107
	st115:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof115
		}
	stCase115:
		if (m.data)[(m.p)] == 45 {
			goto st116
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st139
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st139
			}
		default:
			goto st139
		}
		goto tr107
	st116:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof116
		}
	stCase116:
		if (m.data)[(m.p)] == 45 {
			goto st117
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st138
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st138
			}
		default:
			goto st138
		}
		goto tr107
	st117:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof117
		}
	stCase117:
		if (m.data)[(m.p)] == 45 {
			goto st118
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st137
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st137
			}
		default:
			goto st137
		}
		goto tr107
	st118:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof118
		}
	stCase118:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st119
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st119
			}
		default:
			goto st119
		}
		goto tr107
	st119:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof119
		}
	stCase119:
		if (m.data)[(m.p)] == 58 {
			goto tr169
		}
		goto tr107
	tr169:

		output.ID = string(m.text())

		goto st120
	st120:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof120
		}
	stCase120:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr171
		case 37:
			goto tr172
		case 61:
			goto tr171
		case 95:
			goto tr171
		case 126:
			goto tr171
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto tr171
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr171
				}
			case (m.data)[(m.p)] >= 64:
				goto tr171
			}
		default:
			goto tr171
		}
		goto tr170
	tr171:

		m.pb = m.p

		goto st184
	st184:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof184
		}
	stCase184:
		switch (m.data)[(m.p)] {
		case 33:
			goto st184
		case 35:
			goto tr224
		case 37:
			goto st123
		case 61:
			goto st184
		case 63:
			goto tr226
		case 95:
			goto st184
		case 126:
			goto st184
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st184
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st184
			}
		default:
			goto st184
		}
		goto tr170
	tr224:

		output.SS = string(m.text())
		// Iterate upper letters lowering them
		for _, i := range m.tolower {
			m.data[m.pb+i] = m.data[m.pb+i] + 32
		}
		output.norm = string(m.text())

		goto st185
	tr232:

		output.rComponent = string(m.text())

		goto st185
	tr238:

		if output.rStart {
			m.err = fmt.Errorf(err8141RComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.rStart = true

		output.rComponent = string(m.text())

		goto st185
	tr244:

		output.qComponent = string(m.text())

		goto st185
	tr249:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		output.qComponent = string(m.text())

		goto st185
	st185:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof185
		}
	stCase185:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr227
		case 37:
			goto tr228
		case 61:
			goto tr227
		case 95:
			goto tr227
		case 126:
			goto tr227
		}
		switch {
		case (m.data)[(m.p)] < 63:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto tr227
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr227
			}
		default:
			goto tr227
		}
		goto st0
	tr227:

		m.pb = m.p

		goto st186
	st186:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof186
		}
	stCase186:
		switch (m.data)[(m.p)] {
		case 33:
			goto st186
		case 37:
			goto st121
		case 61:
			goto st186
		case 95:
			goto st186
		case 126:
			goto st186
		}
		switch {
		case (m.data)[(m.p)] < 63:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st186
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st186
			}
		default:
			goto st186
		}
		goto st0
	tr228:

		m.pb = m.p

		goto st121
	st121:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof121
		}
	stCase121:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st122
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st122
			}
		default:
			goto tr175
		}
		goto tr173
	tr175:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st122
	st122:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof122
		}
	stCase122:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st187
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st187
			}
		default:
			goto tr177
		}
		goto tr173
	tr177:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st187
	st187:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof187
		}
	stCase187:
		switch (m.data)[(m.p)] {
		case 33:
			goto st186
		case 37:
			goto st121
		case 61:
			goto st186
		case 95:
			goto st186
		case 126:
			goto st186
		}
		switch {
		case (m.data)[(m.p)] < 63:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st186
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st186
			}
		default:
			goto st186
		}
		goto tr173
	tr172:

		m.pb = m.p

		goto st123
	st123:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof123
		}
	stCase123:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st124
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st124
			}
		default:
			goto tr180
		}
		goto tr178
	tr180:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st124
	st124:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof124
		}
	stCase124:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st188
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st188
			}
		default:
			goto tr182
		}
		goto tr178
	tr182:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st188
	st188:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof188
		}
	stCase188:
		switch (m.data)[(m.p)] {
		case 33:
			goto st184
		case 35:
			goto tr224
		case 37:
			goto st123
		case 61:
			goto st184
		case 63:
			goto tr226
		case 95:
			goto st184
		case 126:
			goto st184
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st184
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st184
			}
		default:
			goto st184
		}
		goto tr178
	tr226:

		output.SS = string(m.text())
		// Iterate upper letters lowering them
		for _, i := range m.tolower {
			m.data[m.pb+i] = m.data[m.pb+i] + 32
		}
		output.norm = string(m.text())

		goto st125
	st125:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof125
		}
	stCase125:
		switch (m.data)[(m.p)] {
		case 43:
			goto st126
		case 61:
			goto st135
		}
		goto st0
	st126:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof126
		}
	stCase126:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr187
		case 37:
			goto tr188
		case 61:
			goto tr187
		case 63:
			goto tr189
		case 95:
			goto tr187
		case 126:
			goto tr187
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto tr187
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr187
				}
			case (m.data)[(m.p)] >= 64:
				goto tr187
			}
		default:
			goto tr187
		}
		goto tr186
	tr187:

		if output.rStart {
			m.err = fmt.Errorf(err8141RComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.rStart = true

		m.pb = m.p

		goto st189
	tr237:

		if output.rStart {
			m.err = fmt.Errorf(err8141RComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.rStart = true

		goto st189
	st189:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof189
		}
	stCase189:
		switch (m.data)[(m.p)] {
		case 33:
			goto st189
		case 35:
			goto tr232
		case 37:
			goto st127
		case 61:
			goto st189
		case 63:
			goto tr234
		case 95:
			goto st189
		case 126:
			goto st189
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st189
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st189
			}
		default:
			goto st189
		}
		goto tr186
	tr188:

		if output.rStart {
			m.err = fmt.Errorf(err8141RComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.rStart = true

		m.pb = m.p

		goto st127
	tr239:

		if output.rStart {
			m.err = fmt.Errorf(err8141RComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.rStart = true

		goto st127
	st127:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof127
		}
	stCase127:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st128
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st128
			}
		default:
			goto tr192
		}
		goto tr190
	tr192:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st128
	st128:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof128
		}
	stCase128:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st190
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st190
			}
		default:
			goto tr194
		}
		goto tr190
	tr194:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st190
	st190:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof190
		}
	stCase190:
		switch (m.data)[(m.p)] {
		case 33:
			goto st189
		case 35:
			goto tr232
		case 37:
			goto st127
		case 61:
			goto st189
		case 63:
			goto tr234
		case 95:
			goto st189
		case 126:
			goto st189
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st189
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st189
			}
		default:
			goto st189
		}
		goto tr190
	tr234:

		output.rComponent = string(m.text())

		goto st191
	tr240:

		if output.rStart {
			m.err = fmt.Errorf(err8141RComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.rStart = true

		output.rComponent = string(m.text())

		goto st191
	st191:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof191
		}
	stCase191:
		switch (m.data)[(m.p)] {
		case 33:
			goto st189
		case 35:
			goto tr232
		case 37:
			goto st127
		case 43:
			goto st192
		case 61:
			goto st193
		case 63:
			goto tr234
		case 95:
			goto st189
		case 126:
			goto st189
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st189
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st189
			}
		default:
			goto st189
		}
		goto tr186
	st192:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof192
		}
	stCase192:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr237
		case 35:
			goto tr238
		case 37:
			goto tr239
		case 47:
			goto st189
		case 61:
			goto tr237
		case 63:
			goto tr240
		case 95:
			goto tr237
		case 126:
			goto tr237
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto tr237
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr237
			}
		default:
			goto tr237
		}
		goto tr186
	st193:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof193
		}
	stCase193:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr205
		case 35:
			goto tr232
		case 37:
			goto tr242
		case 47:
			goto st189
		case 61:
			goto tr205
		case 63:
			goto tr234
		case 95:
			goto tr205
		case 126:
			goto tr205
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto tr205
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr205
			}
		default:
			goto tr205
		}
		goto tr241
	tr205:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		m.pb = m.p

		goto st194
	tr248:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		goto st194
	st194:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof194
		}
	stCase194:
		switch (m.data)[(m.p)] {
		case 33:
			goto st194
		case 35:
			goto tr244
		case 37:
			goto st129
		case 61:
			goto st194
		case 63:
			goto st196
		case 95:
			goto st194
		case 126:
			goto st194
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st194
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st194
			}
		default:
			goto st194
		}
		goto tr204
	tr206:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		m.pb = m.p

		goto st129
	tr250:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		goto st129
	st129:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof129
		}
	stCase129:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st130
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st130
			}
		default:
			goto tr197
		}
		goto tr195
	tr197:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st130
	st130:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof130
		}
	stCase130:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st195
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st195
			}
		default:
			goto tr199
		}
		goto tr195
	tr199:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st195
	st195:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof195
		}
	stCase195:
		switch (m.data)[(m.p)] {
		case 33:
			goto st194
		case 35:
			goto tr244
		case 37:
			goto st129
		case 61:
			goto st194
		case 63:
			goto st196
		case 95:
			goto st194
		case 126:
			goto st194
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st194
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st194
			}
		default:
			goto st194
		}
		goto tr195
	tr251:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		goto st196
	st196:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof196
		}
	stCase196:
		switch (m.data)[(m.p)] {
		case 33:
			goto st194
		case 35:
			goto tr244
		case 37:
			goto st129
		case 61:
			goto st197
		case 63:
			goto st196
		case 95:
			goto st194
		case 126:
			goto st194
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto st194
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st194
			}
		default:
			goto st194
		}
		goto tr204
	st197:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof197
		}
	stCase197:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr248
		case 35:
			goto tr249
		case 37:
			goto tr250
		case 47:
			goto st194
		case 61:
			goto tr248
		case 63:
			goto tr251
		case 95:
			goto tr248
		case 126:
			goto tr248
		}
		switch {
		case (m.data)[(m.p)] < 64:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 59 {
				goto tr248
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr248
			}
		default:
			goto tr248
		}
		goto tr204
	tr242:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		m.pb = m.p

		goto st131
	st131:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof131
		}
	stCase131:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st132
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st132
			}
		default:
			goto tr202
		}
		goto tr200
	tr202:

		// List of positions in the buffer to later lowercase
		m.tolower = append(m.tolower, m.p-m.pb)

		goto st132
	st132:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof132
		}
	stCase132:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st195
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st195
			}
		default:
			goto tr199
		}
		goto tr200
	tr189:

		if output.rStart {
			m.err = fmt.Errorf(err8141RComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.rStart = true

		m.pb = m.p

		goto st133
	st133:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof133
		}
	stCase133:
		if (m.data)[(m.p)] == 43 {
			goto st198
		}
		goto tr186
	st198:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof198
		}
	stCase198:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr237
		case 35:
			goto tr238
		case 37:
			goto tr239
		case 61:
			goto tr237
		case 63:
			goto tr252
		case 95:
			goto tr237
		case 126:
			goto tr237
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto tr237
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr237
				}
			case (m.data)[(m.p)] >= 64:
				goto tr237
			}
		default:
			goto tr237
		}
		goto tr186
	tr252:

		if output.rStart {
			m.err = fmt.Errorf(err8141RComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.rStart = true

		output.rComponent = string(m.text())

		goto st134
	st134:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof134
		}
	stCase134:
		switch (m.data)[(m.p)] {
		case 43:
			goto st198
		case 61:
			goto st135
		}
		goto tr186
	st135:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof135
		}
	stCase135:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr205
		case 37:
			goto tr206
		case 61:
			goto tr205
		case 63:
			goto tr207
		case 95:
			goto tr205
		case 126:
			goto tr205
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto tr205
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr205
				}
			case (m.data)[(m.p)] >= 64:
				goto tr205
			}
		default:
			goto tr205
		}
		goto tr204
	tr207:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		m.pb = m.p

		goto st136
	tr253:

		if output.qStart {
			m.err = fmt.Errorf(err8141QComponentStart, m.p)
			(m.p)--

			{
				goto st201
			}
		}
		output.qStart = true

		goto st136
	st136:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof136
		}
	stCase136:
		if (m.data)[(m.p)] == 61 {
			goto st199
		}
		goto tr204
	st199:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof199
		}
	stCase199:
		switch (m.data)[(m.p)] {
		case 33:
			goto tr248
		case 35:
			goto tr249
		case 37:
			goto tr250
		case 61:
			goto tr248
		case 63:
			goto tr253
		case 95:
			goto tr248
		case 126:
			goto tr248
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 36 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 46 {
				goto tr248
			}
		case (m.data)[(m.p)] > 59:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr248
				}
			case (m.data)[(m.p)] >= 64:
				goto tr248
			}
		default:
			goto tr248
		}
		goto tr204
	st137:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof137
		}
	stCase137:
		if (m.data)[(m.p)] == 58 {
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st119
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st119
			}
		default:
			goto st119
		}
		goto tr107
	st138:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof138
		}
	stCase138:
		switch (m.data)[(m.p)] {
		case 45:
			goto st118
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st137
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st137
			}
		default:
			goto st137
		}
		goto tr107
	st139:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof139
		}
	stCase139:
		switch (m.data)[(m.p)] {
		case 45:
			goto st117
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st138
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st138
			}
		default:
			goto st138
		}
		goto tr107
	st140:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof140
		}
	stCase140:
		switch (m.data)[(m.p)] {
		case 45:
			goto st116
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st139
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st139
			}
		default:
			goto st139
		}
		goto tr107
	st141:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof141
		}
	stCase141:
		switch (m.data)[(m.p)] {
		case 45:
			goto st115
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st140
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st140
			}
		default:
			goto st140
		}
		goto tr107
	st142:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof142
		}
	stCase142:
		switch (m.data)[(m.p)] {
		case 45:
			goto st114
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st141
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st141
			}
		default:
			goto st141
		}
		goto tr107
	st143:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof143
		}
	stCase143:
		switch (m.data)[(m.p)] {
		case 45:
			goto st113
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st142
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st142
			}
		default:
			goto st142
		}
		goto tr107
	st144:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof144
		}
	stCase144:
		switch (m.data)[(m.p)] {
		case 45:
			goto st112
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st143
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st143
			}
		default:
			goto st143
		}
		goto tr107
	st145:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof145
		}
	stCase145:
		switch (m.data)[(m.p)] {
		case 45:
			goto st111
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st144
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st144
			}
		default:
			goto st144
		}
		goto tr107
	st146:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof146
		}
	stCase146:
		switch (m.data)[(m.p)] {
		case 45:
			goto st110
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st145
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st145
			}
		default:
			goto st145
		}
		goto tr107
	st147:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof147
		}
	stCase147:
		switch (m.data)[(m.p)] {
		case 45:
			goto st109
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st146
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st146
			}
		default:
			goto st146
		}
		goto tr107
	st148:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof148
		}
	stCase148:
		switch (m.data)[(m.p)] {
		case 45:
			goto st108
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st147
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st147
			}
		default:
			goto st147
		}
		goto tr107
	st149:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof149
		}
	stCase149:
		switch (m.data)[(m.p)] {
		case 45:
			goto st107
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st148
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto tr107
	st150:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof150
		}
	stCase150:
		switch (m.data)[(m.p)] {
		case 45:
			goto st106
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st149
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st149
			}
		default:
			goto st149
		}
		goto tr107
	st151:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof151
		}
	stCase151:
		switch (m.data)[(m.p)] {
		case 45:
			goto st105
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st150
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st150
			}
		default:
			goto st150
		}
		goto tr107
	st152:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof152
		}
	stCase152:
		switch (m.data)[(m.p)] {
		case 45:
			goto st104
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st151
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st151
			}
		default:
			goto st151
		}
		goto tr107
	st153:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof153
		}
	stCase153:
		switch (m.data)[(m.p)] {
		case 45:
			goto st103
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st152
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st152
			}
		default:
			goto st152
		}
		goto tr107
	st154:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof154
		}
	stCase154:
		switch (m.data)[(m.p)] {
		case 45:
			goto st102
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st153
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st153
			}
		default:
			goto st153
		}
		goto tr107
	st155:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof155
		}
	stCase155:
		switch (m.data)[(m.p)] {
		case 45:
			goto st101
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st154
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st154
			}
		default:
			goto st154
		}
		goto tr107
	st156:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof156
		}
	stCase156:
		switch (m.data)[(m.p)] {
		case 45:
			goto st100
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st155
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st155
			}
		default:
			goto st155
		}
		goto tr107
	st157:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof157
		}
	stCase157:
		switch (m.data)[(m.p)] {
		case 45:
			goto st99
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st156
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st156
			}
		default:
			goto st156
		}
		goto tr107
	st158:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof158
		}
	stCase158:
		switch (m.data)[(m.p)] {
		case 45:
			goto st98
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st157
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st157
			}
		default:
			goto st157
		}
		goto tr107
	st159:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof159
		}
	stCase159:
		switch (m.data)[(m.p)] {
		case 45:
			goto st97
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st158
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st158
			}
		default:
			goto st158
		}
		goto tr107
	st160:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof160
		}
	stCase160:
		switch (m.data)[(m.p)] {
		case 45:
			goto st96
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st159
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st159
			}
		default:
			goto st159
		}
		goto tr107
	st161:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof161
		}
	stCase161:
		switch (m.data)[(m.p)] {
		case 45:
			goto st95
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st160
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st160
			}
		default:
			goto st160
		}
		goto tr107
	st162:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof162
		}
	stCase162:
		switch (m.data)[(m.p)] {
		case 45:
			goto st94
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st161
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st161
			}
		default:
			goto st161
		}
		goto tr107
	st163:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof163
		}
	stCase163:
		switch (m.data)[(m.p)] {
		case 45:
			goto st93
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st162
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st162
			}
		default:
			goto st162
		}
		goto tr107
	st164:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof164
		}
	stCase164:
		switch (m.data)[(m.p)] {
		case 45:
			goto st92
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st163
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st163
			}
		default:
			goto st163
		}
		goto tr107
	st165:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof165
		}
	stCase165:
		switch (m.data)[(m.p)] {
		case 45:
			goto st91
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st164
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st164
			}
		default:
			goto st164
		}
		goto tr107
	st166:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof166
		}
	stCase166:
		switch (m.data)[(m.p)] {
		case 45:
			goto st90
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st165
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st165
			}
		default:
			goto st165
		}
		goto tr107
	tr106:

		m.pb = m.p

		if m.parsingMode != RFC8141Only {
			// Throw an error when:
			// - we are entering here matching the the prefix in the namespace identifier part
			// - looking ahead (3 chars) we find a colon
			if pos := m.p + 3; pos < m.pe && m.data[pos] == 58 && output.prefix != "" {
				m.err = fmt.Errorf(errNoUrnWithinID, pos)
				(m.p)--

				{
					goto st201
				}
			}
		}

		goto st167
	st167:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof167
		}
	stCase167:
		switch (m.data)[(m.p)] {
		case 45:
			goto st89
		case 82:
			goto st168
		case 114:
			goto st168
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st166
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st166
			}
		default:
			goto st166
		}
		goto tr104
	st168:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof168
		}
	stCase168:
		switch (m.data)[(m.p)] {
		case 45:
			goto st90
		case 58:
			goto tr169
		case 78:
			goto st169
		case 110:
			goto st169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st165
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st165
			}
		default:
			goto st165
		}
		goto tr104
	st169:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof169
		}
	stCase169:
		switch (m.data)[(m.p)] {
		case 45:
			goto tr211
		case 58:
			goto tr169
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st164
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st164
			}
		default:
			goto st164
		}
		goto tr107
	tr211:

		output.prefix = string(m.text())

		goto st170
	st170:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof170
		}
	stCase170:
		switch (m.data)[(m.p)] {
		case 45:
			goto st92
		case 48:
			goto st171
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st163
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st171
			}
		default:
			goto st171
		}
		goto tr107
	st171:

		(m.p)--

		m.err = fmt.Errorf(err8141InformalID, m.p)
		{
			goto st201
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof171
		}
	stCase171:
		if (m.data)[(m.p)] == 45 {
			goto st93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st162
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st162
			}
		default:
			goto st162
		}
		goto tr107
	stCase172:
		switch (m.data)[(m.p)] {
		case 85:
			goto tr213
		case 117:
			goto tr213
		}
		goto tr52
	tr213:

		m.pb = m.p

		if m.parsingMode != RFC8141Only {
			// Throw an error when:
			// - we are entering here matching the the prefix in the namespace identifier part
			// - looking ahead (3 chars) we find a colon
			if pos := m.p + 3; pos < m.pe && m.data[pos] == 58 && output.prefix != "" {
				m.err = fmt.Errorf(errNoUrnWithinID, pos)
				(m.p)--

				{
					goto st201
				}
			}
		}

		goto st173
	st173:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof173
		}
	stCase173:
		switch (m.data)[(m.p)] {
		case 82:
			goto st174
		case 114:
			goto st174
		}
		goto tr52
	st174:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof174
		}
	stCase174:
		switch (m.data)[(m.p)] {
		case 78:
			goto st175
		case 110:
			goto st175
		}
		goto tr52
	st175:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof175
		}
	stCase175:
		if (m.data)[(m.p)] == 58 {
			goto tr216
		}
		goto tr52
	tr216:

		output.prefix = string(m.text())
		{
			goto st87
		}
		goto st200
	st200:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof200
		}
	stCase200:
		goto tr52
	st201:
		if (m.p)++; (m.p) == (m.pe) {
			goto _testEof201
		}
	stCase201:
		switch (m.data)[(m.p)] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st201
	stOut:
	_testEof2:
		m.cs = 2
		goto _testEof
	_testEof3:
		m.cs = 3
		goto _testEof
	_testEof4:
		m.cs = 4
		goto _testEof
	_testEof176:
		m.cs = 176
		goto _testEof
	_testEof5:
		m.cs = 5
		goto _testEof
	_testEof6:
		m.cs = 6
		goto _testEof
	_testEof7:
		m.cs = 7
		goto _testEof
	_testEof8:
		m.cs = 8
		goto _testEof
	_testEof9:
		m.cs = 9
		goto _testEof
	_testEof10:
		m.cs = 10
		goto _testEof
	_testEof11:
		m.cs = 11
		goto _testEof
	_testEof12:
		m.cs = 12
		goto _testEof
	_testEof13:
		m.cs = 13
		goto _testEof
	_testEof14:
		m.cs = 14
		goto _testEof
	_testEof15:
		m.cs = 15
		goto _testEof
	_testEof16:
		m.cs = 16
		goto _testEof
	_testEof17:
		m.cs = 17
		goto _testEof
	_testEof18:
		m.cs = 18
		goto _testEof
	_testEof19:
		m.cs = 19
		goto _testEof
	_testEof20:
		m.cs = 20
		goto _testEof
	_testEof21:
		m.cs = 21
		goto _testEof
	_testEof22:
		m.cs = 22
		goto _testEof
	_testEof23:
		m.cs = 23
		goto _testEof
	_testEof24:
		m.cs = 24
		goto _testEof
	_testEof25:
		m.cs = 25
		goto _testEof
	_testEof26:
		m.cs = 26
		goto _testEof
	_testEof27:
		m.cs = 27
		goto _testEof
	_testEof28:
		m.cs = 28
		goto _testEof
	_testEof29:
		m.cs = 29
		goto _testEof
	_testEof30:
		m.cs = 30
		goto _testEof
	_testEof31:
		m.cs = 31
		goto _testEof
	_testEof32:
		m.cs = 32
		goto _testEof
	_testEof33:
		m.cs = 33
		goto _testEof
	_testEof34:
		m.cs = 34
		goto _testEof
	_testEof35:
		m.cs = 35
		goto _testEof
	_testEof36:
		m.cs = 36
		goto _testEof
	_testEof37:
		m.cs = 37
		goto _testEof
	_testEof38:
		m.cs = 38
		goto _testEof
	_testEof177:
		m.cs = 177
		goto _testEof
	_testEof39:
		m.cs = 39
		goto _testEof
	_testEof40:
		m.cs = 40
		goto _testEof
	_testEof178:
		m.cs = 178
		goto _testEof
	_testEof41:
		m.cs = 41
		goto _testEof
	_testEof42:
		m.cs = 42
		goto _testEof
	_testEof43:
		m.cs = 43
		goto _testEof
	_testEof45:
		m.cs = 45
		goto _testEof
	_testEof46:
		m.cs = 46
		goto _testEof
	_testEof47:
		m.cs = 47
		goto _testEof
	_testEof179:
		m.cs = 179
		goto _testEof
	_testEof48:
		m.cs = 48
		goto _testEof
	_testEof49:
		m.cs = 49
		goto _testEof
	_testEof50:
		m.cs = 50
		goto _testEof
	_testEof51:
		m.cs = 51
		goto _testEof
	_testEof52:
		m.cs = 52
		goto _testEof
	_testEof53:
		m.cs = 53
		goto _testEof
	_testEof54:
		m.cs = 54
		goto _testEof
	_testEof55:
		m.cs = 55
		goto _testEof
	_testEof56:
		m.cs = 56
		goto _testEof
	_testEof57:
		m.cs = 57
		goto _testEof
	_testEof58:
		m.cs = 58
		goto _testEof
	_testEof59:
		m.cs = 59
		goto _testEof
	_testEof60:
		m.cs = 60
		goto _testEof
	_testEof61:
		m.cs = 61
		goto _testEof
	_testEof62:
		m.cs = 62
		goto _testEof
	_testEof63:
		m.cs = 63
		goto _testEof
	_testEof64:
		m.cs = 64
		goto _testEof
	_testEof65:
		m.cs = 65
		goto _testEof
	_testEof66:
		m.cs = 66
		goto _testEof
	_testEof67:
		m.cs = 67
		goto _testEof
	_testEof68:
		m.cs = 68
		goto _testEof
	_testEof69:
		m.cs = 69
		goto _testEof
	_testEof180:
		m.cs = 180
		goto _testEof
	_testEof70:
		m.cs = 70
		goto _testEof
	_testEof181:
		m.cs = 181
		goto _testEof
	_testEof71:
		m.cs = 71
		goto _testEof
	_testEof72:
		m.cs = 72
		goto _testEof
	_testEof182:
		m.cs = 182
		goto _testEof
	_testEof73:
		m.cs = 73
		goto _testEof
	_testEof74:
		m.cs = 74
		goto _testEof
	_testEof75:
		m.cs = 75
		goto _testEof
	_testEof76:
		m.cs = 76
		goto _testEof
	_testEof77:
		m.cs = 77
		goto _testEof
	_testEof78:
		m.cs = 78
		goto _testEof
	_testEof79:
		m.cs = 79
		goto _testEof
	_testEof80:
		m.cs = 80
		goto _testEof
	_testEof81:
		m.cs = 81
		goto _testEof
	_testEof82:
		m.cs = 82
		goto _testEof
	_testEof84:
		m.cs = 84
		goto _testEof
	_testEof85:
		m.cs = 85
		goto _testEof
	_testEof86:
		m.cs = 86
		goto _testEof
	_testEof183:
		m.cs = 183
		goto _testEof
	_testEof87:
		m.cs = 87
		goto _testEof
	_testEof88:
		m.cs = 88
		goto _testEof
	_testEof89:
		m.cs = 89
		goto _testEof
	_testEof90:
		m.cs = 90
		goto _testEof
	_testEof91:
		m.cs = 91
		goto _testEof
	_testEof92:
		m.cs = 92
		goto _testEof
	_testEof93:
		m.cs = 93
		goto _testEof
	_testEof94:
		m.cs = 94
		goto _testEof
	_testEof95:
		m.cs = 95
		goto _testEof
	_testEof96:
		m.cs = 96
		goto _testEof
	_testEof97:
		m.cs = 97
		goto _testEof
	_testEof98:
		m.cs = 98
		goto _testEof
	_testEof99:
		m.cs = 99
		goto _testEof
	_testEof100:
		m.cs = 100
		goto _testEof
	_testEof101:
		m.cs = 101
		goto _testEof
	_testEof102:
		m.cs = 102
		goto _testEof
	_testEof103:
		m.cs = 103
		goto _testEof
	_testEof104:
		m.cs = 104
		goto _testEof
	_testEof105:
		m.cs = 105
		goto _testEof
	_testEof106:
		m.cs = 106
		goto _testEof
	_testEof107:
		m.cs = 107
		goto _testEof
	_testEof108:
		m.cs = 108
		goto _testEof
	_testEof109:
		m.cs = 109
		goto _testEof
	_testEof110:
		m.cs = 110
		goto _testEof
	_testEof111:
		m.cs = 111
		goto _testEof
	_testEof112:
		m.cs = 112
		goto _testEof
	_testEof113:
		m.cs = 113
		goto _testEof
	_testEof114:
		m.cs = 114
		goto _testEof
	_testEof115:
		m.cs = 115
		goto _testEof
	_testEof116:
		m.cs = 116
		goto _testEof
	_testEof117:
		m.cs = 117
		goto _testEof
	_testEof118:
		m.cs = 118
		goto _testEof
	_testEof119:
		m.cs = 119
		goto _testEof
	_testEof120:
		m.cs = 120
		goto _testEof
	_testEof184:
		m.cs = 184
		goto _testEof
	_testEof185:
		m.cs = 185
		goto _testEof
	_testEof186:
		m.cs = 186
		goto _testEof
	_testEof121:
		m.cs = 121
		goto _testEof
	_testEof122:
		m.cs = 122
		goto _testEof
	_testEof187:
		m.cs = 187
		goto _testEof
	_testEof123:
		m.cs = 123
		goto _testEof
	_testEof124:
		m.cs = 124
		goto _testEof
	_testEof188:
		m.cs = 188
		goto _testEof
	_testEof125:
		m.cs = 125
		goto _testEof
	_testEof126:
		m.cs = 126
		goto _testEof
	_testEof189:
		m.cs = 189
		goto _testEof
	_testEof127:
		m.cs = 127
		goto _testEof
	_testEof128:
		m.cs = 128
		goto _testEof
	_testEof190:
		m.cs = 190
		goto _testEof
	_testEof191:
		m.cs = 191
		goto _testEof
	_testEof192:
		m.cs = 192
		goto _testEof
	_testEof193:
		m.cs = 193
		goto _testEof
	_testEof194:
		m.cs = 194
		goto _testEof
	_testEof129:
		m.cs = 129
		goto _testEof
	_testEof130:
		m.cs = 130
		goto _testEof
	_testEof195:
		m.cs = 195
		goto _testEof
	_testEof196:
		m.cs = 196
		goto _testEof
	_testEof197:
		m.cs = 197
		goto _testEof
	_testEof131:
		m.cs = 131
		goto _testEof
	_testEof132:
		m.cs = 132
		goto _testEof
	_testEof133:
		m.cs = 133
		goto _testEof
	_testEof198:
		m.cs = 198
		goto _testEof
	_testEof134:
		m.cs = 134
		goto _testEof
	_testEof135:
		m.cs = 135
		goto _testEof
	_testEof136:
		m.cs = 136
		goto _testEof
	_testEof199:
		m.cs = 199
		goto _testEof
	_testEof137:
		m.cs = 137
		goto _testEof
	_testEof138:
		m.cs = 138
		goto _testEof
	_testEof139:
		m.cs = 139
		goto _testEof
	_testEof140:
		m.cs = 140
		goto _testEof
	_testEof141:
		m.cs = 141
		goto _testEof
	_testEof142:
		m.cs = 142
		goto _testEof
	_testEof143:
		m.cs = 143
		goto _testEof
	_testEof144:
		m.cs = 144
		goto _testEof
	_testEof145:
		m.cs = 145
		goto _testEof
	_testEof146:
		m.cs = 146
		goto _testEof
	_testEof147:
		m.cs = 147
		goto _testEof
	_testEof148:
		m.cs = 148
		goto _testEof
	_testEof149:
		m.cs = 149
		goto _testEof
	_testEof150:
		m.cs = 150
		goto _testEof
	_testEof151:
		m.cs = 151
		goto _testEof
	_testEof152:
		m.cs = 152
		goto _testEof
	_testEof153:
		m.cs = 153
		goto _testEof
	_testEof154:
		m.cs = 154
		goto _testEof
	_testEof155:
		m.cs = 155
		goto _testEof
	_testEof156:
		m.cs = 156
		goto _testEof
	_testEof157:
		m.cs = 157
		goto _testEof
	_testEof158:
		m.cs = 158
		goto _testEof
	_testEof159:
		m.cs = 159
		goto _testEof
	_testEof160:
		m.cs = 160
		goto _testEof
	_testEof161:
		m.cs = 161
		goto _testEof
	_testEof162:
		m.cs = 162
		goto _testEof
	_testEof163:
		m.cs = 163
		goto _testEof
	_testEof164:
		m.cs = 164
		goto _testEof
	_testEof165:
		m.cs = 165
		goto _testEof
	_testEof166:
		m.cs = 166
		goto _testEof
	_testEof167:
		m.cs = 167
		goto _testEof
	_testEof168:
		m.cs = 168
		goto _testEof
	_testEof169:
		m.cs = 169
		goto _testEof
	_testEof170:
		m.cs = 170
		goto _testEof
	_testEof171:
		m.cs = 171
		goto _testEof
	_testEof173:
		m.cs = 173
		goto _testEof
	_testEof174:
		m.cs = 174
		goto _testEof
	_testEof175:
		m.cs = 175
		goto _testEof
	_testEof200:
		m.cs = 200
		goto _testEof
	_testEof201:
		m.cs = 201
		goto _testEof

	_testEof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 44, 45, 46, 47, 83, 84, 85, 86, 172, 173, 174, 175:

				m.err = fmt.Errorf(errPrefix, m.p)
				(m.p)--

				{
					goto st201
				}

			case 121, 122:

				if m.parsingMode == RFC2141Only || m.parsingMode == All {
					m.err = fmt.Errorf(errHex, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				// Otherwise, we expect the machine to fallback to SCIM errors

			case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64:

				// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
				if m.parsingMode == All {
					// TODO: store why the machine fallback to the RFC2141 one?
					output.scim = nil
					// Rewind the cursor after the prefix ends ("urn:")
					(m.p) = (4) - 1

					// Go to the "urn" machine from this point on
					{
						goto st5
					}
				}
				m.err = fmt.Errorf(errSCIMNamespace, m.p)
				(m.p)--

				{
					goto st201
				}

			case 65, 66, 67, 68, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82:

				// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
				if m.parsingMode == All {
					// TODO: store why the machine fallback to the RFC2141 one?
					output.scim = nil
					// Rewind the cursor after the prefix ends ("urn:")
					(m.p) = (4) - 1

					// Go to the "urn" machine from this point on
					{
						goto st5
					}
				}
				m.err = fmt.Errorf(errSCIMType, m.p)
				(m.p)--

				{
					goto st201
				}

			case 69:

				// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
				if m.parsingMode == All {
					// TODO: store why the machine fallback to the RFC2141 one?
					output.scim = nil
					// Rewind the cursor after the prefix ends ("urn:")
					(m.p) = (4) - 1

					// Go to the "urn" machine from this point on
					{
						goto st5
					}
				}
				m.err = fmt.Errorf(errSCIMName, m.p)
				(m.p)--

				{
					goto st201
				}

			case 70:

				// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
				if m.parsingMode == All {
					// TODO: store why the machine fallback to the RFC2141 one?
					output.scim = nil
					// Rewind the cursor after the prefix ends ("urn:")
					(m.p) = (4) - 1

					// Go to the "urn" machine from this point on
					{
						goto st5
					}
				}
				if m.p == m.pe {
					m.err = fmt.Errorf(errSCIMOtherIncomplete, m.p-1)
				} else {
					m.err = fmt.Errorf(errSCIMOther, m.p)
				}
				(m.p)--

				{
					goto st201
				}

			case 120:

				m.err = fmt.Errorf(err8141SpecificString, m.p)
				(m.p)--

				{
					goto st201
				}

			case 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 169, 170, 171:

				m.err = fmt.Errorf(err8141Identifier, m.p)
				(m.p)--

				{
					goto st201
				}

			case 126, 133, 134:

				m.err = fmt.Errorf(err8141MalformedRComp, m.p)
				(m.p)--

				{
					goto st201
				}

			case 135, 136:

				m.err = fmt.Errorf(err8141MalformedQComp, m.p)
				(m.p)--

				{
					goto st201
				}

			case 177, 178:

				output.SS = string(m.text())
				// Iterate upper letters lowering them
				for _, i := range m.tolower {
					m.data[m.pb+i] = m.data[m.pb+i] + 32
				}
				output.norm = string(m.text())

				output.kind = RFC2141

			case 184, 188:

				output.SS = string(m.text())
				// Iterate upper letters lowering them
				for _, i := range m.tolower {
					m.data[m.pb+i] = m.data[m.pb+i] + 32
				}
				output.norm = string(m.text())

				output.kind = RFC8141

			case 1, 2, 3, 4:

				m.err = fmt.Errorf(errPrefix, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errParse, m.p)
				(m.p)--

				{
					goto st201
				}

			case 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:

				m.err = fmt.Errorf(errIdentifier, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errParse, m.p)
				(m.p)--

				{
					goto st201
				}

			case 38:

				m.err = fmt.Errorf(errSpecificString, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errParse, m.p)
				(m.p)--

				{
					goto st201
				}

			case 71, 72:

				if m.parsingMode == RFC2141Only || m.parsingMode == All {
					m.err = fmt.Errorf(errHex, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				// Otherwise, we expect the machine to fallback to SCIM errors

				// In case we are in fallback mode we are now gonna jump to normal RFC2141 URN parsing
				if m.parsingMode == All {
					// TODO: store why the machine fallback to the RFC2141 one?
					output.scim = nil
					// Rewind the cursor after the prefix ends ("urn:")
					(m.p) = (4) - 1

					// Go to the "urn" machine from this point on
					{
						goto st5
					}
				}
				if m.p == m.pe {
					m.err = fmt.Errorf(errSCIMOtherIncomplete, m.p-1)
				} else {
					m.err = fmt.Errorf(errSCIMOther, m.p)
				}
				(m.p)--

				{
					goto st201
				}

			case 123, 124:

				if m.parsingMode == RFC2141Only || m.parsingMode == All {
					m.err = fmt.Errorf(errHex, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				// Otherwise, we expect the machine to fallback to SCIM errors

				m.err = fmt.Errorf(err8141SpecificString, m.p)
				(m.p)--

				{
					goto st201
				}

			case 127, 128:

				if m.parsingMode == RFC2141Only || m.parsingMode == All {
					m.err = fmt.Errorf(errHex, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				// Otherwise, we expect the machine to fallback to SCIM errors

				m.err = fmt.Errorf(err8141MalformedRComp, m.p)
				(m.p)--

				{
					goto st201
				}

			case 129, 130:

				if m.parsingMode == RFC2141Only || m.parsingMode == All {
					m.err = fmt.Errorf(errHex, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				// Otherwise, we expect the machine to fallback to SCIM errors

				m.err = fmt.Errorf(err8141MalformedQComp, m.p)
				(m.p)--

				{
					goto st201
				}

			case 87, 167, 168:

				m.err = fmt.Errorf(err8141Identifier, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errPrefix, m.p)
				(m.p)--

				{
					goto st201
				}

			case 189, 190, 191, 193:

				output.rComponent = string(m.text())

				output.kind = RFC8141

			case 194, 195, 196:

				output.qComponent = string(m.text())

				output.kind = RFC8141

			case 186, 187:

				output.fComponent = string(m.text())

				output.kind = RFC8141

			case 185:

				m.pb = m.p

				output.fComponent = string(m.text())

				output.kind = RFC8141

			case 5, 41, 42:

				m.err = fmt.Errorf(errIdentifier, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errPrefix, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errParse, m.p)
				(m.p)--

				{
					goto st201
				}

			case 43:

				m.err = fmt.Errorf(errIdentifier, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errNoUrnWithinID, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errParse, m.p)
				(m.p)--

				{
					goto st201
				}

			case 39, 40:

				if m.parsingMode == RFC2141Only || m.parsingMode == All {
					m.err = fmt.Errorf(errHex, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				// Otherwise, we expect the machine to fallback to SCIM errors

				m.err = fmt.Errorf(errSpecificString, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(errParse, m.p)
				(m.p)--

				{
					goto st201
				}

			case 131, 132:

				if m.parsingMode == RFC2141Only || m.parsingMode == All {
					m.err = fmt.Errorf(errHex, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				// Otherwise, we expect the machine to fallback to SCIM errors

				m.err = fmt.Errorf(err8141MalformedRComp, m.p)
				(m.p)--

				{
					goto st201
				}

				m.err = fmt.Errorf(err8141MalformedQComp, m.p)
				(m.p)--

				{
					goto st201
				}

			case 180:

				output.scim.Name = string(m.data[output.scim.pos:m.p])

				output.SS = string(m.text())
				// Iterate upper letters lowering them
				for _, i := range m.tolower {
					m.data[m.pb+i] = m.data[m.pb+i] + 32
				}
				output.norm = string(m.text())

				output.kind = RFC7643

			case 181, 182:

				output.scim.Other = string(m.data[output.scim.pos:m.p])

				output.SS = string(m.text())
				// Iterate upper letters lowering them
				for _, i := range m.tolower {
					m.data[m.pb+i] = m.data[m.pb+i] + 32
				}
				output.norm = string(m.text())

				output.kind = RFC7643

			case 192, 198:

				if output.rStart {
					m.err = fmt.Errorf(err8141RComponentStart, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				output.rStart = true

				output.rComponent = string(m.text())

				output.kind = RFC8141

			case 197, 199:

				if output.qStart {
					m.err = fmt.Errorf(err8141QComponentStart, m.p)
					(m.p)--

					{
						goto st201
					}
				}
				output.qStart = true

				output.qComponent = string(m.text())

				output.kind = RFC8141
			}
		}

	_out:
		{
		}
	}

	if m.cs < firstFinal || m.cs == enFail {
		return nil, m.err
	}

	return output, nil
}

func (m *machine) WithParsingMode(x ParsingMode) {
	m.parsingMode = x
	m.parsingModeSet = true
}
