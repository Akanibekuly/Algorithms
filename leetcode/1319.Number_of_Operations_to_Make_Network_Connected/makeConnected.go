package Number_of_Operations_to_Make_Network_Connected

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	adj := make([][]int, n)
	for _, c := range connections {
		adj[c[0]] = append(adj[c[0]], c[1])
		adj[c[1]] = append(adj[c[1]], c[0])
	}

	visit := make([]bool, n)
	var dfs func(int)
	dfs = func(node int) {
		visit[node] = true
		for _, neighbor := range adj[node] {
			if !visit[neighbor] {
				dfs(neighbor)
			}
		}
	}

	var numberOfConnectedComponents int
	for i := 0; i < n; i++ {
		if !visit[i] {
			numberOfConnectedComponents++
			dfs(i)
		}
	}

	return numberOfConnectedComponents - 1
}
