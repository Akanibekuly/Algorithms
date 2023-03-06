package passthepillow

func passThePillow(n int, time int) int {
	rem := (time)%(2*n-2) + 1

	if rem <= n {
		return rem
	}

	return 2*n - 1 - rem
}
