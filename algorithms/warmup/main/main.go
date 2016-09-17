package main

import (
	"fmt"
	"hackerrank/algorithms/warmup"
)

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = warmup.SolveMeFirst(a, b)
	fmt.Println(res)
}
