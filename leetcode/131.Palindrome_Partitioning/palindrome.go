package palindromepartitioning

func partition(s string) [][]string {

	var res [][]string
	for i := 1; i <= len(s); i++ {
		tmp := s[:i]
		if isPolindrome(tmp) {
			// fmt.Println(tmp, fmt.Sprintf("[%s]", s[i:]))
			if s[i:] == "" {
				res = append(res, []string{tmp})
			} else {
				arr := partition(s[i:])
				for i := range arr {
					res = append(res, append([]string{tmp}, arr[i]...))
				}
			}
		}
	}
	return res
}

func isPolindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
