package excelsheetcolumntitle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSheeet(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  string
	}{
		{"1", 1, "A"},
		{"28", 28, "AB"},
		{"701->ZY", 701, "ZY"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, convertToTitle(c.n))
		})
	}
}
