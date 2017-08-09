package divisible_sum_pairs

// https://www.hackerrank.com/challenges/divisible-sum-pairs

// 输入元素个数 n ， 目标除数 k
// 输出 a0 ～ an-1 两数相加能被 k 整除的数目

func divisible_sum_pairs(arr []int, tar int) (count int) {
	count = 0
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if (arr[i]+arr[j])%tar == 0 {
				count++
			}
		}
	}
	return
}
