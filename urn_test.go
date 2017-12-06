package urn

import (
	"fmt"
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

func TestLexicalEquivalence(t *testing.T) {
	for ii, tt := range equivalenceTests {
		urnlx, oklx := Parse(tt.lx)
		urnrx, okrx := Parse(tt.rx)

		if oklx && okrx {

			require.True(t, urnlx.Equal(urnlx))
			require.True(t, urnrx.Equal(urnrx))

			if tt.eq {
				require.True(t, urnlx.Equal(urnrx), ierror(ii))
				require.True(t, urnrx.Equal(urnlx), ierror(ii))
			} else {
				require.False(t, urnlx.Equal(urnrx), ierror(ii))
				require.False(t, urnrx.Equal(urnlx), ierror(ii))
			}
		} else {
			t.Log("Something wrong in the testing table ...")
		}
	}
}

func TestDefaultPrefixWhenString(t *testing.T) {
	u := &URN{
		ID: "pippo",
		SS: "pluto",
	}

	require.Equal(t, "urn:pippo:pluto", u.String())
}

func BenchmarkParse(b *testing.B) {
	for ii, tt := range tests {
		outcome := (map[bool]string{true: "ok", false: "no"})[tt.ok]
		b.Run(
			fmt.Sprintf("%s/%02d/%s%0*s/", outcome, ii, tt.in, len(tt.in)-80, " "),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Parse(tt.in)
				}
			},
		)
	}
}
