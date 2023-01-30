package lfucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLFUCache1(t *testing.T) {
	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)

	assert.Equal(t, 1, lfu.Get(1))

	lfu.Put(3, 3)

	assert.Equal(t, -1, lfu.Get(2))
	assert.Equal(t, 3, lfu.Get(3))

	lfu.Put(4, 4)

	assert.Equal(t, -1, lfu.Get(1))
	assert.Equal(t, 3, lfu.Get(3))
	assert.Equal(t, 4, lfu.Get(4))
}

func TestLFUcCacheZeroCacpacity(t *testing.T) {
	lfu := Constructor(0)

	assert.Equal(t, -1, lfu.Get(0))

	lfu.Put(0, 0)

	assert.Equal(t, -1, lfu.Get(0))
}

func TestLFUCache2(t *testing.T) {
	lfu := Constructor(2)

	assert := assert.New(t)
	assert.Equal(-1, lfu.Get(2))

	lfu.Put(2, 6)

	assert.Equal(-1, lfu.Get(1))

	lfu.Put(1, 5)
	lfu.Put(1, 2)

	assert.Equal(2, lfu.Get(1))
	assert.Equal(6, lfu.Get(2))
}

func TestLFUCache3(t *testing.T) {
	lfu := Constructor(2)

	lfu.Put(2, 1)
	lfu.Put(3, 2)

	assert := assert.New(t)
	assert.Equal(2, lfu.Get(3))
	assert.Equal(1, lfu.Get(2))

	lfu.Put(4, 3)

	assert.Equal(1, lfu.Get(2))
	assert.Equal(-1, lfu.Get(3))
	assert.Equal(3, lfu.Get(4))
}

// [2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
// ,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]
func TestLFUCache(t *testing.T) {

	cases := []struct {
		name string
		cmd  []string
		args [][]int
		res  []int
	}{
		{
			name: "1",
			cmd:  []string{"LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"},
			args: [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}, {4, 4}, {1}, {3}, {4}},
			res:  []int{-2, -2, -2, 1, -2, -1, 3, -2, -1, 3, 4},
		},
		{
			cmd:  []string{"LFUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"},
			args: [][]int{{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}},
			res:  []int{-2, -2, -2, -2, -2, -2, -1, -2, 19, 17, -2, -1, -2, -2, -2, -1, -2, -1, 5, -1, 12, -2, -2, 3, 5, 5, -2, -2, 1, -2, -1, -2, 30, 5, 30, -2, -2, -2, -1, -2, -1, 24, -2, -2, 18, -2, -2, -2, -2, 14, -2, -2, 18, -2, -2, 11, -2, -2, -2, -2, -2, 18, -2, -2, -1, -2, 4, 29, 30, -2, 12, 11, -2, -2, -2, -2, 29, -2, -2, -2, -2, 17, -1, 18, -2, -2, -2, -1, -2, -2, -2, 20, -2, -2, -2, 29, 18, 18, -2, -2, -2, -2, 20, -2, -2, -2, -2, -2, -2, -2},
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert := assert.New(tt)
			if !(assert.Equal(len(c.args), len(c.cmd)) && assert.Equal(len(c.cmd), len(c.res))) {
				tt.Logf("invalid len cmd [%d] and args [%d] and res [%d]", len(c.cmd), len(c.args), len(c.res))
				tt.FailNow()
			}

			var l LFUCachef
			for j := range c.cmd {
				switch c.cmd[j] {
				case "LFUCache":
					require.Equal(tt, len(c.args[j]), 1, "case: %d cmd: %d", i+1, j+1)
					l = *Constructorf(c.args[j][0])
				case "put":
					require.Equal(tt, len(c.args[j]), 2, "case: %d cmd: %d", i+1, j+1)
					// t.Logf("Put key [%d] value [%d]", c.args[j][0], c.args[j][1])
					l.Put(c.args[j][0], c.args[j][1])
					// t.Logf("after put len: %d", len(l.cache))
				case "get":
					require.Equal(tt, len(c.args[j]), 1, "case: %d cmd: %d", i+1, j+1)
					assert.Equal(c.res[j], l.Get(c.args[j][0]), "case: %d cmd: %d key: %d", i+1, j+1, c.args[j][0])
				}
			}
		})
	}
}
