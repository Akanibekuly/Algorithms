package main

type treeNode struct {
	Val     int
	IsApple bool
	Childs  []*treeNode
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	treeMap := make(map[int]*treeNode, n)
	root := &treeNode{
		Val:     0,
		IsApple: hasApple[0],
	}
	treeMap[0] = root

	for _, edge := range edges {
		// fmt.Println(edge)
		var parent, child *treeNode
		for _, vertex := range edge {
			node, ok := treeMap[vertex]
			if ok {
				parent = node
			} else {
				child = &treeNode{
					Val:     vertex,
					IsApple: hasApple[vertex],
				}
				treeMap[vertex] = child
			}
		}
		// fmt.Println(*parent, *child)
		parent.Childs = append(parent.Childs, child)
	}

	count := 0
	for _, child := range root.Childs {
		count += traverse(child)
	}

	return count
}

func traverse(node *treeNode) int {
	if node == nil {
		return 0
	}

	count := 0
	for _, child := range node.Childs {
		count += traverse(child)
	}

	if count != 0 || node.IsApple {
		count += 2
	}

	return count
}

func minTimeDFS(n int, edges [][]int, hasApple []bool) int {

	graph := make([][]int, n)

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	time := make([]int, n)

	var dfs func(u, par int)
	dfs = func(u, par int) {
		for _, v := range graph[u] {
			if v != par {
				dfs(v, u)
				if hasApple[v] {
					hasApple[u] = true
					time[u] += time[v] + 2
				}
			}
		}
	}

	dfs(0, -1)
	return time[0]
}
