package leetcode

func minFlipsMonoIncr(s string) int {
	total := 0
	for i := range s {
		if s[i] == '1' {
			total++
		}
	}

	if total == 0 || total == len(s) {
		return 0
	}

	count, min, ln := 0, total, len(s)
	if min > ln-total {
		min = ln - total
	}

	for i := range s {
		if s[i] == '1' {
			count++
		}

		tmp := ln + 2*count - (i + 1 + total)
		if tmp <= min {
			min = tmp
		}

		// min == 0, "000 111111"
		// count == total "001101011 000000000"
		if min == 0 || count == total {
			return min
		}
	}

	return min
}
