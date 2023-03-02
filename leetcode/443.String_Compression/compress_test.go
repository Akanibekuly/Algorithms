package stringcompression

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	cases := []struct {
		chars []byte
		after []byte
		res   int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '3'}, 6},
		{[]byte{'a'}, []byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, []byte{'a', 'b', '1', '2'}, 4},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, compress(c.chars))
			for j := 0; j < c.res; j++ {
				assert.Equal(tt, c.after[j], c.chars[j], "")
			}
		})
	}
}
