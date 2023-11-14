package uniquelength3palindromicsubsequences

func countPalindromicSubsequence(s string) int {
	first, last := [26]int{}, [26]int{}
	for i := range first {
		first[i] = -1
	}
	for i := range s {
		curr := s[i] - 'a'
		if first[curr] == -1 {
			first[curr] = i
		}

		last[curr] = i
	}

	var ans int
	for i := 0; i < 26; i++ {
		if first[i] == -1 {
			continue
		}

		between := make(map[byte]struct{})
		for j := first[i] + 1; j < last[i]; j++ {
			between[s[j]] = struct{}{}
		}

		ans += len(between)
	}
	return ans
}
