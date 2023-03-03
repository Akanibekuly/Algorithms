package str

import "fmt"

func strStr(haystack string, needle string) int {
	fmt.Println(haystack, needle)
	if len(needle) > len(haystack) {
		return -1
	}

	if haystack == needle {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
