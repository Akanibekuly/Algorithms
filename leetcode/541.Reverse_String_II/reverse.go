package reversestringii

import (
	"strings"
)

func reverseStr(s string, k int) string {
	builder := strings.Builder{}

	var i int
	for i < len(s) {
		if len(s) < i+k {
			for j := len(s) - 1; j >= i; j-- {
				builder.WriteByte(s[j])
			}
			return builder.String()
		} else {
			for j := i + k - 1; j >= i; j-- {
				builder.WriteByte(s[j])
			}
			i += k
			if i+k < len(s) {
				builder.WriteString(s[i : i+k])
				i += k
			} else {
				builder.WriteString(s[i:])
				break
			}
		}
	}

	return builder.String()
}

func reverseStr2(s string, k int) string {
	if k == 1 {
		return s
	}

	res := []byte(s)
	for i := 0; i < len(s); i++ {
		if i%(2*k) != 0 {
			continue
		}

		l, r := i, i+k-1
		if r >= len(s) {
			r = len(s) - 1
		}
		i = r
		for l < r {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		}
	}

	return string(res)
}
