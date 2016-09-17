package main

import (
	"fmt"
	"hackerrank/tutorials"
)

func main() {
	var primes []int
	tutorials.ComputePrime(&primes, 10000000)
	fmt.Println(len(primes))
	for _, v := range primes {
		fmt.Println(v)
	}
}
