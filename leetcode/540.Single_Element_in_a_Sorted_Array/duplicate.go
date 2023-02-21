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

func singleNonDuplicate2(nums []int) int {
	l, r := 0, len(nums)-1
	if len(nums) == 1 {
		return nums[0]
	}

	for l <= r {
		m := (l + r) >> 1

		isEven := m%2 == 0

		if l == r {
			return nums[l]
		}

		if nums[m] != nums[m-1] && nums[m] != nums[m+1] {
			return nums[m]
		}

		if m > 0 && nums[m] == nums[m-1] {
			if isEven {
				r = m - 2
			} else {
				l = m + 1
			}
		} else if m < len(nums)-1 && nums[m] == nums[m+1] {
			if isEven {
				l = m + 2
			} else {
				r = m - 1
			}
		}
	}

	return nums[l]
}
