package numberofflowersinfullbloom

import (
	"fmt"
	"sort"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	start, end := make(sort.IntSlice, len(flowers)), make(sort.IntSlice, len(flowers))
	for i, v := range flowers {
		start[i], end[i] = v[0], v[1]
	}
	start.Sort()
	end.Sort()

	ans := make([]int, len(people))
	for i, v := range people {
		ans[i] = start.Search(v+1) - end.Search(v)
	}

	return ans
}

func fullBloomFlowersWithDiffs(flowers [][]int, people []int) []int {

	diffs := make(map[int]int)
	for i := range flowers {
		diffs[flowers[i][0]]++
		diffs[flowers[i][1]]--
	}

	difference := make([][]int, 0, len(diffs))
	for k, v := range diffs {
		difference = append(difference, []int{k, v})
	}
	sort.Slice(difference, func(i, j int) bool {
		return difference[i][0] < difference[j][0]
	})
	fmt.Println(difference)

	prev := 0
	for i := range difference {
		difference[i][1] += prev
		prev = difference[i][1]
	}

	fmt.Println(difference)
	binarySearch := func(target int) int {
		l, r := 0, len(difference)
		for l < r {
			mid := (l + r) / 2
			if target < difference[mid][0] {
				r = mid
			} else {
				l = mid + 1
			}
		}

		return l
	}

	ans := make([]int, len(people))
	for j := range people {
		i := binarySearch(people[j]) - 1
		ans[j] = difference[i][1]
	}

	return ans
}
