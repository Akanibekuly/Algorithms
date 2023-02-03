package zigzagconversion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		s, res string
		n      int
	}{
		{s: "PAYPALISHIRING", n: 3, res: "PAHNAPLSIIGYIR"},
		{s: "PAYPALISHIRING", n: 4, res: "PINALSIGYAHRPI"},
		{s: "A", n: 1, res: "A"},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, convert(c.s, c.n))
		})
	}
}
