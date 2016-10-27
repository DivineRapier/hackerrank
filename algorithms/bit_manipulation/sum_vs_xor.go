package bit_manipulation

import "math"

// 异或操作：
// 0 ^ 0 = 0
// 0 ^ 1 = 1
// 1 ^ 0 = 1
// 1 ^ 1 = 0
// 所以，可以发现 0 + a = 0 ^ a

func sum_vs_xor(n int) (count int64) {

	count = int64(math.Pow(2.0, float64(zero_count(n))))

	return
}

func zero_count(n int) (count int) {

	for c := 1; c <= n; c <<= 1 {
		if c&n == 0 {
			count++
		}
	}
	return
}
