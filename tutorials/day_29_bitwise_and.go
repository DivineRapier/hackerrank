package tutorials

import "fmt"

// bitwise_and
// 输入两个数 n, k
// 对于区间 (0, n] 寻找两个数  A, B  A & B < k 的最大值

func getMax(n, k int) int {
	max := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if r := i & j; r < k && r > max {
				max = r
			}
		}
	}
	return max
}

func bitwise_and() {
	var t, n, k int
	fmt.Scan(&t)

	for a := 0; a < t; a++ {
		fmt.Scan(&n)
		fmt.Scan(&k)
		fmt.Println(getMax(n, k))
	}

}
