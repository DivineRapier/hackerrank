package main

import "fmt"

func lonely_integer(nums []int) int {
	lnum := 0
	for _, v := range nums {
		lnum ^= v
	}
	return lnum
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	num := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num[i])
	}
	a := lonely_integer(num)
	fmt.Printf("%d\n", a)
}
