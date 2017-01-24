package main

import "math"

// https://www.hackerrank.com/challenges/is-fibo

// Fn = (a^n - b^n) / sqrt(5)
// a = (1 + sqrt(5))/2
// b = (1 - sqrt(5))/2

// a = sqrt(5 * n * n + 4);
// b = sqrt(5 * n * n - 4);

func is_fibo(n int64) bool {
	a := math.Sqrt(float64(5*n*n + 4))
	b := math.Sqrt(float64(5*n*n - 4))
	return a == float64(int64(a)) || b == float64(int64(b))
}
