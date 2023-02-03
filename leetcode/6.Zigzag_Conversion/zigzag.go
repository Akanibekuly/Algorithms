package zigzagconversion

func convert(s string, n int) string {
	if n == 1 {
		return s
	}

	res := make([]byte, 0, len(s))
	// move by row
	for i := 0; i < n; i++ {
		// by cycle 2*n-2
		for k := i; k < len(s); k += 2*n - 2 {
			res = append(res, s[k])
			// check if the next fits
			if i != 0 && i != n-1 && k+2*n-2-i*2 < len(s) {
				res = append(res, s[k+2*n-2-i*2])
			}
		}
	}

	return string(res)
}
