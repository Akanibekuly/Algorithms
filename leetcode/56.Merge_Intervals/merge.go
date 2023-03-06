package mergeintervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0
	for i < len(intervals)-1 {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1])
			intervals = append(intervals[:i+1], intervals[i+2:]...)
		} else {
			i++
		}
	}

	return intervals
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
