package reorganizestring

func reorganizeString(s string) string {
	freq := make([]int, 26)
	for _, r := range s {
		freq[int(r-'a')]++
	}

	max, letter := -1, -1
	for i := range freq {
		if freq[i] > max {
			max = freq[i]
			letter = i
		}
	}

	if max > (len(s)+1)/2 {
		return ""
	}
	var index int
	arr := make([]rune, len(s))
	for freq[letter] > 0 {
		arr[index] = rune('a' + letter)
		index += 2
		freq[letter]--
	}

	for i := range freq {
		for freq[i] > 0 {
			if index >= len(s) {
				index = 1
			}
			arr[index] = rune('a' + i)
			index += 2
			freq[i]--
		}
	}

	return string(arr)
}
