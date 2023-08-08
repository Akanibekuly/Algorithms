package searchinrotatedsortedarray

func search(nums []int, target int) int {
	return binSearch(nums, target, findPivot(nums))
}

func findPivot(nums []int) int {
	l, r := 0, len(nums)-1
	for l != r {
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func binSearch(nums []int, target, pivot int) int {
	l, r := 0, len(nums)-1
	for l != r {
		m := (l + r) / 2
		if nums[(m+pivot)%len(nums)] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	if nums[(l+pivot)%len(nums)] != target {
		return -1
	}

	return (l + pivot) % len(nums)
}

func searchOld(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}

		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] > nums[right] {
			left += 1
			right -= 1
		} else {
			if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

	}

	return -1
}
