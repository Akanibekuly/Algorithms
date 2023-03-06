package passthepillow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPass(t *testing.T) {
	cases := []struct {
		n, time, res int
	}{
		{4, 5, 2},
		{3, 2, 3},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.res, passThePillow(c.n, c.time))
		})
	}
}
