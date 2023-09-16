package extracharactersinastring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtraCharacters(t *testing.T) {
	cases := []struct {
		name, s string
		dict    []string
		res     int
	}{}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, minExtraChar(c.s, c.dict))
		})
	}
}
