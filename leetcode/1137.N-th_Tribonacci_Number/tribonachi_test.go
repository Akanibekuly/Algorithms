package nthtribonaccinumber

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTribonachi(t *testing.T) {
	cases := []struct {
		n, res int
	}{
		{4, 4}, {25, 1389537}, {34, 334745777},
	}

	for i, c := range cases {
		start := time.Now()
		t.Run(fmt.Sprintf("case: %d, n: %d", i+1, c.n), func(tt *testing.T) {
			assert.Equal(tt, c.res, tribonacci(c.n))
		})

		t.Log("time: ", time.Since(start))
	}
}
