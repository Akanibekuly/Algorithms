package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	s, dict := [26]int{}, [26]int{}
	for i := range s1 {
		s[s1[i]-'a']++
		dict[s2[i]-'a']++
	}

	if s == dict {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		dict[s2[i]-'a']++
		dict[s2[i-len(s1)]-'a']--
		if dict == s {
			return true
		}
	}

	return false
}
