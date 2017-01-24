package main

import "fmt"

// DynamicArray ...
func DynamicArray() {
	N, Q, lastAns := 0, 0, 0
	var a, x, y int
	fmt.Scanf("%d%d", &N, &Q)
	lists := make([][]int, N, N)
	for i := 0; i < N; i++ {
		lists[i] = make([]int, 0)
	}
	for i := 0; i < Q; i++ {
		fmt.Scanf("%d%d%d", &a, &x, &y)
		switch a {
		case 1:
			lists[(x^lastAns)%N] = append(lists[(x^lastAns)%N], y)
		case 2:
			lastAns = lists[(x^lastAns)%N][y%len(lists[(x^lastAns)%N])]
			fmt.Println(lastAns)
		}
	}
}

func main() {
	DynamicArray()
}
