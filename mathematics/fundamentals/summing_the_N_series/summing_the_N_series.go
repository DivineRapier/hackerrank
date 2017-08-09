package summing_the_N_series

import "fmt"

// M = 10 ^ 9 + 7
// Tn = n * n - (n - 1) * (n - 1)
// Tn = (n + n - 1) * (n - n + 1) = 2n - 1
// Sn = (T1 + T2 + T3 + ... + Tn) mod M
//    = (1 + 3 + 5 + ... + (2n - 1)) mod M
//    = ((1 + (2n - 1)) * n) /2 mod M
//    = n*n mod M
//    = (n mod M) * (n mod M)

const M = 1000000007

func compute(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = ((arr[i] % M) * (arr[i] % M)) % M
	}
}

func print(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
