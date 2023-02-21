package duplicate

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)/2
	for l < r {
		m := (l + r) / 2
		if nums[2*m] == nums[2*m+1] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l*2]
}
