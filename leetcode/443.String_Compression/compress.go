package stringcompression

import (
	"strconv"
)

func compress(chars []byte) int {
	var count, idx int
	curr := chars[0]
	for i := 0; i < len(chars); i++ {
		if chars[i] == curr {
			count++
		} else {
			chars[idx], idx = curr, idx+1
			if count != 1 {
				s := []byte(strconv.Itoa(count))
				for j := range s {
					chars[idx], idx = s[j], idx+1
				}
			}

			curr = chars[i]
			count = 1
		}
	}

	chars[idx], idx = curr, idx+1
	if count != 1 {
		s := []byte(strconv.Itoa(count))
		for j := range s {
			chars[idx], idx = s[j], idx+1
		}
	}

	return idx
}
