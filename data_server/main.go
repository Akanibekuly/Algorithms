package main

import (
	"fmt"
	"strconv"
)

var ServerMap = map[int][]int{}
var visited = map[int]bool{}
var claster = map[int]int{}

func main() {
	var n int
	fmt.Scan(&n)
	// clean map and visited table for routes
	ServerMap = map[int][]int{}
	visited = map[int]bool{}
	// load server names
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if _, k := ServerMap[x]; !k {
			ServerMap[x] = []int{y}
			visited[x] = false
		} else {
			ServerMap[x] = append(ServerMap[x], y)
		}
		if _, k := ServerMap[y]; !k {
			ServerMap[y] = []int{x}
			visited[y] = false
		} else {
			ServerMap[y] = append(ServerMap[y], x)
		}
	}
	// fmt.Println(ServerMap)
	c := 1
	var servCount int
	fmt.Scan(&servCount)
	res := ""
	for i := 0; i < servCount; i++ {
		var ServerName, count int
		fmt.Scan(&ServerName, &count)
		tres := ""
		cnt := 0
		if !visited[ServerName] {
			dfs(ServerName, c)
			c++
		}
		for j := 0; j < count; j++ {
			var temp int
			fmt.Scan(&temp)

			if claster[temp] != 0 && (claster[ServerName] == claster[temp]) {
				cnt++
				tres += " " + strconv.Itoa(temp)
			}
		}

		if cnt != 0 {
			tres = strconv.Itoa(cnt) + tres + "\n"
		} else {
			tres = "0\n"
		}
		res += tres
	}
	fmt.Println(claster)
	fmt.Print(res)
}

func dfs(curr, c int) {
	visited[curr] = true
	claster[curr] = c
	for _, v := range ServerMap[curr] {
		if !visited[v] {
			dfs(v, c)
		}
	}
}
