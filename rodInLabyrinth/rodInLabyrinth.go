package main

func main() {}

type Rod struct {
	X   int
	Y   int
	Pos string
}

func rodInLabyrinth(labyrinth [][]string) int {
	count := 0
	rod=Rod{
		X: 0,
		y: 1,
		pos: "h",
	}
	
	return count
}

func isPossible(move string, lab [][]string, rod Rod) bool {


