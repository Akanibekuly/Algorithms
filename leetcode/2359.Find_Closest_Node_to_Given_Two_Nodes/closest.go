package closest

const MAX int = 1e6

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	distance1, distance2 := make([]int, n), make([]int, n)
	solve(edges, distance1, node1)
	solve(edges, distance2, node2)

	ans, maxDist := n, MAX
	for i := 0; i < n; i++ {
		dist := distance1[i]
		if distance2[i] > dist {
			dist = distance2[i]
		}
		if dist < maxDist {
			maxDist = dist
			ans = i
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

func solve(edges []int, distance []int, node int) {
	for i := 0; i < len(distance); i++ {
		distance[i] = MAX
	}
	distance[node] = 0
	for edges[node] != -1 && distance[edges[node]] == MAX {
		distance[edges[node]] = distance[node] + 1
		node = edges[node]
	}
}
