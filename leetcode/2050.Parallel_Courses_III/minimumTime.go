package parallelcoursesiii

import "fmt"

func minimumTime(n int, relations [][]int, time []int) int {
	graph := make(map[int][]int, n)
	indegree := make([]int, n)
	for _, edge := range relations {
		x, y := edge[0]-1, edge[1]-1
		graph[x] = append(graph[x], y)
		indegree[y]++
	}
	// fmt.Println(graph, indegree)

	queue, maxTime := &Queue{}, make([]int, n)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue.add(i)
			maxTime[i] = time[i]
			// fmt.Print(i, " ")
			// queue.print()
		}
	}
	// queue.print()

	for !queue.isEmpty() {
		node := queue.remove()
		for _, neighbour := range graph[node] {
			maxTime[neighbour] = max(maxTime[neighbour], maxTime[node]+time[neighbour])
			indegree[neighbour]--
			if indegree[neighbour] == 0 {
				queue.add(neighbour)
			}
		}

		// fmt.Println("after", node, graph[node], indegree, queue.isEmpty())
		// queue.print()
	}

	var ans int
	for i := range maxTime {
		ans = max(ans, maxTime[i])
	}

	return ans
}

type listNode struct {
	Value int
	Next  *listNode
}

type Queue struct {
	head, tail *listNode
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

func (q *Queue) remove() int {
	if q.isEmpty() {
		return -1
	}

	val := q.head.Value
	q.head = q.head.Next
	if q.head == nil {
		q.tail = q.head
	}
	return val
}

func (q *Queue) add(n int) {
	tmp := &listNode{Value: n}
	if q.tail == nil {
		q.head = tmp
		q.tail = q.head
		return
	}

	q.tail.Next = tmp
	q.tail = q.tail.Next
}

func (q *Queue) print() {
	if q.head == nil {
		fmt.Println("<nil>")
		return
	}

	for t := q.head; t != nil; t = t.Next {
		fmt.Print(t.Value)
		if t.Next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
