package gcd

import "strings"

// recirsive
func gcdOfStringsRecursive(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		return gcdOfStringsRecursive(str2, str1)
	}

	if str1[:len(str2)] != str2 {
		return ""
	}

	// режем по кусочкам
	for len(str1) >= len(str2) && str1[:len(str2)] == str2 {
		str1 = str1[len(str2):]
	}

	if str1 == "" {
		return str2
	}

	return gcdOfStringsRecursive(str2, str1)
}

func gcdOfStrings(str1, str2 string) string {
	long, short := str1, str2

	if long < short {
		long, short = short, long
	}

	var curr string
	for i := len(short) - 1; i >= 0; i-- {
		curr = string([]byte(short)[:i+1])
		if strings.Count(long, curr)*(i+1) == len(long) && strings.Count(short, curr)*(i+1) == len(short) {
			return curr
		}
	}

	return ""
}
