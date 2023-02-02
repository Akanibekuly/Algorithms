package verifyinganaliendictionary

func isAlienSorted(words []string, order string) bool {
	dict := make(map[byte]int, len(order))
	for i := range order {
		dict[order[i]] = i
	}

	isAlienOrder := func(a, b string) bool {
		for i := 0; i < len(a) && i < len(b); i++ {
			if dict[a[i]] > dict[b[i]] {
				return false
			} else if dict[a[i]] < dict[b[i]] {
				return true
			}
		}

		return len(a) <= len(b)
	}

	for i := 0; i < len(words)-1; i++ {
		if !isAlienOrder(words[i], words[i+1]) {
			return false
		}
	}

	return true
}
