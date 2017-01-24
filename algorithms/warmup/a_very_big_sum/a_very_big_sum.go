package main

// https://www.hackerrank.com/challenges/a-very-big-sum?h_r=next-challenge&h_v=zen

func a_very_big_sum(arr []int64) int64 {
	var ret int64
	for _, v := range arr {
		ret += v
	}
	return ret
}
