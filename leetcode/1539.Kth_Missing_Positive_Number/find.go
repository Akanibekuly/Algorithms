package kthmissingpositivenumber

func findKthPositive(arr []int, k int) int {
	// lower
	if k < arr[0] {
		return k
	}

	// higher
	if k > arr[len(arr)-1]-len(arr) {
		return k + len(arr)
	}

	// find position in the array
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		miss := arr[m] - len(arr[:m]) - 1
		// fmt.Printf("l %d, r %d, m  %d, miss %d,  k %d, len(arr[:m]) %d\n", l, r, m, miss, k, len(arr[:m]))

		if k <= miss {
			r = m
		} else {
			l = m + 1
		}

		if l == r {
			// fmt.Printf("[l==r] l %d, r %d, m  %d, miss %d,  k %d, len(arr[:m]) %d\n", l, r, m, miss, k, len(arr[:m]))
			return k + len(arr[:l])
		}
	}

	return -1
}
