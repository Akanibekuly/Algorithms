package maximalnetworkrank

func maximalNetworkRank(n int, roads [][]int) int {
	graph := make([][]bool, n)

	for i := range graph {
		graph[i] = make([]bool, n)
	}

	degree := make([]int, n)

	for _, r := range roads {
		graph[r[0]][r[1]] = true
		graph[r[1]][r[0]] = true

		degree[r[0]]++
		degree[r[1]]++
	}

	var max int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			netw := degree[i] + degree[j]
			if graph[i][j] {
				netw--
			}
			if max < netw {
				max = netw
			}
		}
	}

	return max
}

func maximalNetworkRankHash(n int, roads [][]int) int {
	cities := make([]map[int]struct{}, n)
	for i := range cities {
		cities[i] = make(map[int]struct{})
	}

	for _, r := range roads {
		cities[r[0]][r[1]] = struct{}{}
		cities[r[1]][r[0]] = struct{}{}
	}

	max := -1
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			netw := len(cities[i]) + len(cities[j])
			if _, ok := cities[i][j]; ok {
				netw--
			}

			if netw > max {
				max = netw
			}
		}
	}

	return max
}
