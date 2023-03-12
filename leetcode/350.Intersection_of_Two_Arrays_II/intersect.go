package intersectionoftwoarraysii

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	set1, set2 := make(map[int]int), make(map[int]int)
	for i := range nums1 {
		set1[nums1[i]]++
	}
	for i := range nums2 {
		set2[nums2[i]]++
	}

	var res []int

	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	for n1, c1 := range set1 {
		if c2, exists := set2[n1]; exists {
			size := c1
			if c2 < c1 {
				size = c2
			}
			tmp := make([]int, size)
			for i := range tmp {
				tmp[i] = n1
			}
			res = append(res, tmp...)
		}
	}
	return res
}
