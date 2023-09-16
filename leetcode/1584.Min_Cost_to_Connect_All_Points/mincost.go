package mincosttoconnectallpoints

import (
	"math"
	"sort"
)

type path struct {
	dist, p1, p2 int
}

func minCostConnectPointsUnionFind(points [][]int) int {
	parents := make([]int, len(points))
	for i := range parents {
		parents[i] = i
	}

	var find func(x int) int
	var union func(x, y int) bool
	var getDist func(p1, p2 int) int

	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	union = func(x, y int) bool {
		r1, r2 := find(x), find(y)
		if r1 != r2 {
			parents[r2] = r1
			return true
		}
		return false
	}

	getDist = func(p1, p2 int) int {
		return abs(points[p1][0]-points[p2][0]) + abs(points[p1][1]-points[p2][1])
	}

	paths := make([]path, 0, len(points)*(len(points)-1)/2)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			paths = append(paths, path{getDist(i, j), i, j})
		}
	}

	sort.Slice(paths, func(i, j int) bool {
		return paths[i].dist < paths[j].dist
	})

	ans := 0
	for _, p := range paths {
		if union(p.p1, p.p2) {
			ans += p.dist
		}
	}

	return ans
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func minCostConnectPoints(points [][]int) int {
	dp := make([]int, len(points))
	seen := make([]bool, len(points))
	for i := range dp {
		dp[i] = math.MaxInt
	}

	var dist, curr, best int
	var p1, p2 []int

	for i := 1; i < len(points); i++ {
		best = curr
		dp[curr] = math.MaxInt
		seen[curr] = true

		for j := 0; j < len(points); j++ {
			if seen[j] {
				continue
			}
			p1, p2 = points[curr], points[j]
			dp[j] = min(dp[j], abs(p1[0]-p2[0])+abs(p1[1]-p2[1]))
			if dp[j] < dp[best] {
				best = j
			}
		}
		dist += dp[best]
		curr = best
	}

	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
