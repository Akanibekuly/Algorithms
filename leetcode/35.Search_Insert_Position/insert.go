package searchinsertposition

func searchInsert(nums []int, target int) int {

	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}

	if l > 0 && nums[l-1] == target {
		return l - 1
	}

	return r
}
