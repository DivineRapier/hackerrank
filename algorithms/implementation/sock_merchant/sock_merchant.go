package main

import "sort"

func sock_merchant(number []int) int {
	if number == nil {
		return -1
	}
	cnt := 0

	sort.Ints(number)
	length := len(number)
	for i := 0; i < length; {
		if i == length-1 {
			break
		}
		if number[i] == number[i+1] {
			cnt++
			i += 2
		} else {
			i++
		}
	}

	return cnt
}
