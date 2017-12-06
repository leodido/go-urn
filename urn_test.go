package urn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUrn(t *testing.T) {
	for ii, tt := range tests {
		urn, ok := Parse(tt.in)

		if ok {
			require.True(t, tt.ok, herror(ii, tt))
			require.Equal(t, tt.obj.prefix, urn.prefix, herror(ii, tt))
			require.Equal(t, tt.obj.ID, urn.ID, herror(ii, tt))
			require.Equal(t, tt.obj.SS, urn.SS, herror(ii, tt))
			require.Equal(t, tt.str, urn.String(), herror(ii, tt))
			require.Equal(t, tt.norm, urn.Normalize().String(), herror(ii, tt))
		} else {
			require.False(t, tt.ok, herror(ii, tt))
			require.Empty(t, urn, herror(ii, tt))
		}
	}
}
