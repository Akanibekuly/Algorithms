package minimumpenaltyforashop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBestClosingTimee(t *testing.T) {
	cases := []struct {
		customers string
		res       int
	}{
		{"YYNY", 2},
		{"NNNNN", 0},
		{"YYYY", 4},
	}

	for _, c := range cases {
		t.Run(c.customers, func(t *testing.T) {
			require.Equal(t, c.res, bestClosingTime(c.customers))
		})
	}
}
