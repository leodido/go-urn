package urn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
			// require.Empty(t, urn.Tree(), herror(ii, tt))
		}
	}
}
