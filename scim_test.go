package urn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalParseAsSCIM(t *testing.T) {
	tests := []struct {
		name     string
		in       []byte
		expected *SCIM
	}{
		{
			"simple scim sub-namespace with schemas type",
			[]byte("urn:ietf:params:scim:schemas:core:2.0:User"),
			&SCIM{
				urn: &URN{
					prefix: "urn",
					ID:     "ietf",
					SS:     "params:scim:schemas:core:2.0:User",
					norm:   "params:scim:schemas:core:2.0:User",
				},
				Type:  "schemas",
				Name:  "core",
				Other: "2.0:User",
			},
		},
		{
			"complex scim sub-namespace with schemas type",
			[]byte("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified"),
			&SCIM{
				urn: &URN{
					prefix: "urn",
					ID:     "ietf",
					SS:     "params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified",
					norm:   "params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified",
				},
				Type:  "schemas",
				Name:  "extension",
				Other: "enterprise:2.0:User:meta.lastModified",
			},
		},
		{
			"simple scim sub-namespace with api type",
			[]byte("urn:ietf:params:scim:api:messages:2.0:ListResponse"),
			&SCIM{
				urn: &URN{
					prefix: "urn",
					ID:     "ietf",
					SS:     "params:scim:api:messages:2.0:ListResponse",
					norm:   "params:scim:api:messages:2.0:ListResponse",
				},
				Type:  "api",
				Name:  "messages",
				Other: "2.0:ListResponse",
			},
		},
	}
	for _, tt := range tests {
		actual, err := ParseAsSCIM(tt.in)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, tt.expected, actual)
	}
}

func TestErrorParseAsSCIM(t *testing.T) {
	tests := []struct {
		name     string
		in       []byte
		expected string
	}{
		{
			"invalid urn string",
			[]byte("urn-ietf-params"),
			"failed to parse: invalid URN",
		},
		{
			"invalid scim sub-namespace",
			[]byte("urn:ietf:param:ext-scim:schemas:core:2.0:User"),
			"faled to parse: invalid SCIM Sub-namespace",
		},

		{
			"invalid scim sub-namespace",
			[]byte("urn:ietf:params:scim:model:core:2.0:User"),
			`got invalid type in SCIM Sub-namespace, either "schemas" or "api": model`,
		},
	}
	for _, tt := range tests {
		_, err := ParseAsSCIM(tt.in)
		assert.Error(t, err, tt.in)
		if tt.expected != err.Error() {
			t.Errorf("want '%s' as an error message, but got '%s'",
				tt.expected, err.Error())
		}
	}
}
