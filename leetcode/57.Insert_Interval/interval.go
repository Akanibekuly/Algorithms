package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := [][]int{}
	start := 0
	for start < len(intervals) && intervals[start][0] < newInterval[0] {
		result = append(result, intervals[start])
		start++
	}

	if start == 0 || intervals[start-1][1] < newInterval[0] {
		result = append(result, newInterval)
	} else {
		intervals[start-1][1] = max(intervals[start-1][1], newInterval[1])
	}

	prev := result[len(result)-1]
	for i := start; i < len(intervals); i++ {
		current := intervals[i]
		if prev[1] >= current[0] {
			prev[1] = max(prev[1], current[1])
		} else {
			result = append(result, current)
			prev = current
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// first approach
// func insert(intervals [][]int, newInterval []int) [][]int {

// 	start_idx, start_pos := 0, newInterval[0]
// 	for start_idx < len(intervals) && intervals[start_idx][0] <= start_pos {
// 		if intervals[start_idx][1] >= start_pos {
// 			start_pos = intervals[start_idx][0]
// 			break
// 		}

// 		start_idx++
// 	}

// 	last_idx, last_pos := start_idx, newInterval[1]
// 	for last_idx < len(intervals) && intervals[last_idx][0] <= last_pos {
// 		if last_pos <= intervals[last_idx][1] {
// 			last_pos = intervals[last_idx][1]
// 			last_idx++
// 			break
// 		}
// 		last_idx++
// 	}

// 	// fmt.Println(start_idx, last_idx)
// 	var res [][]int
// 	res = append(res, intervals[:start_idx]...)
// 	res = append(res, []int{start_pos, last_pos})
// 	res = append(res, intervals[last_idx:]...)
// 	return res
// }
