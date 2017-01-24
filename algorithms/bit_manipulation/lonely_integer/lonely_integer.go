package main

func lonely_integer(nums []int) int {
	lnum := 0
	for _, v := range nums {
		lnum ^= v
	}
	return lnum
}
