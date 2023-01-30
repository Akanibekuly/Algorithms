package nthtribonaccinumber

var cache = map[int]int{0: 0, 1: 1, 2: 1}

// too slow
func tribonacciRec(n int) int {
	if res, ok := cache[n]; ok {
		return res
	}

	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	if n < 1 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	var d int
	for i := 3; i <= n; i++ {
		d = a + b + c
		a, b, c = b, c, d
	}

	return d
}
