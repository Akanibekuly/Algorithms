package findthearrayconcatenationvalue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	cases := []struct {
		nums []int
		ans  int64
	}{
		{[]int{7, 52, 2, 4}, 696},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.ans, findTheArrayConcVal(c.nums))
		})
	}
}
