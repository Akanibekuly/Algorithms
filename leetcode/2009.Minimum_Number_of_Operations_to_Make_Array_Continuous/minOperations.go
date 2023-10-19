package minimumnumberofoperationstomakearraycontinuous

import "sort"

func minOperations(nums []int) int {
	dict := make(map[int]struct{})
	for i := range nums {
		dict[nums[i]] = struct{}{}
	}

	arr, n := make([]int, 0, len(dict)), len(nums)
	for k := range dict {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	ans := n
	for i := range arr {
		count := binarySearch(arr, arr[i]+n-1) - i
		if ans < n-count {
			ans = n - count
		}
	}

	return ans
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
