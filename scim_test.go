package urn

import (
	"encoding/json"
	"testing"

	scimschema "github.com/leodido/go-urn/scim/schema"
	"github.com/stretchr/testify/require"
)

func TestSCIMJSONMarshaling(t *testing.T) {
	t.Run("roundtrip", func(t *testing.T) {
		// Marshal
		exp := SCIM{Type: scimschema.Schemas, Name: "core", Other: "extension:enterprise:2.0:User"}
		mar, err := json.Marshal(exp)
		require.NoError(t, err)

		// Unmarshal
		var act SCIM
		err = json.Unmarshal(mar, &act)
		require.NoError(t, err)

		require.Equal(t, exp, act)
	})
}
