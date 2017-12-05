package urn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUrnParse(t *testing.T) {
	for ii, tt := range tests {
		urn, ok := Parse(tt.in)

		if ok {
			require.True(t, tt.ok, herror(ii, tt))
			require.Equal(t, tt.obj.ID, urn.ID, herror(ii, tt))
			require.Equal(t, tt.obj.SS, urn.SS, herror(ii, tt))
			// (todo) > test normalized version vs expected one
		} else {
			require.False(t, tt.ok, herror(ii, tt))
			require.Empty(t, urn, herror(ii, tt))
		}
	}
}
