package warmup

import "math"

// https://www.hackerrank.com/challenges/diagonal-difference?h_r=next-challenge&h_v=zen
func diagonal_difference(matrix [][]int) float64 {
	sum1, sum2 := 0, 0
	l := len(matrix)
	for i := 0; i < l; i++ {
		sum1 += matrix[i][i]
		sum2 += matrix[i][l-i-1]
	}
	return math.Abs(float64(sum1 - sum2))
}
