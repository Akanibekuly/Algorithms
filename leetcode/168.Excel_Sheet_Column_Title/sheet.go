package excelsheetcolumntitle

func convertToTitle(n int) string {
	var arr []rune
	for n > 0 {
		n -= 1
		arr = append(arr, 'A'+rune(n%26))
		n /= 26
	}

	// reverse
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return string(arr)
}
