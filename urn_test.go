package urn

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultPrefixWhenString(t *testing.T) {
	u := &URN{
		ID: "pippo",
		SS: "pluto",
	}

	assert.Equal(t, "urn:pippo:pluto", u.String())
}

func TestParseSignature(t *testing.T) {
	urn, ok := Parse([]byte(``))
	assert.Nil(t, urn)
	assert.False(t, ok)
}

func TestJSONMarshaling(t *testing.T) {
	t.Run("roundtrip", func(t *testing.T) {
		// Marshal
		expected := URN{ID: "oid", SS: "1.2.3.4"}
		bytes, err := json.Marshal(expected)
		if !assert.NoError(t, err) {
			return
		}
		// Unmarshal
		var actual URN
		err = json.Unmarshal(bytes, &actual)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, expected.String(), actual.String())
	})

	t.Run("invalid URN", func(t *testing.T) {
		var actual URN
		err := json.Unmarshal([]byte(`"not a URN"`), &actual)
		assert.EqualError(t, err, "invalid URN: not a URN")
	})

	t.Run("empty", func(t *testing.T) {
		var actual URN
		err := actual.UnmarshalJSON(nil)
		assert.EqualError(t, err, "unexpected end of JSON input")
	})
}

func TestPercentEncodingRequiresHexDigits(t *testing.T) {
	for _, in := range []string{
		"urn:example:foo%2g",
		"urn:example:foo%2G",
		"urn:example:foo%gg",
	} {
		_, ok := Parse([]byte(in))
		assert.False(t, ok, in)
		_, ok = Parse([]byte(in), WithParsingMode(RFC8141Only))
		assert.False(t, ok, "8141 "+in)
	}

	u, ok := Parse([]byte("urn:example:foo%2Fbar"))
	assert.True(t, ok)
	assert.Equal(t, "foo%2Fbar", u.SS)
	assert.Equal(t, "foo%2fbar", u.norm)

	u, ok = Parse([]byte("URN:x:abc%1Dz%2F%3az"))
	assert.True(t, ok)
	assert.Equal(t, "abc%1Dz%2F%3az", u.SS)
	assert.Equal(t, "abc%1dz%2f%3az", u.norm)
}
