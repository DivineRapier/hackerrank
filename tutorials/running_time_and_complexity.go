package tutorials

// ComputePrime ...
func ComputePrime(primes *[]int, rag int) {
	*primes = append(*primes, 2)
	*primes = append(*primes, 3)

	for i := 5; i < rag; i++ {
		j := 0
		for ; j < len(*primes); j++ {
			if i%((*primes)[j]) == 0 {
				break
			}
		}
		if j == len(*primes) {
			*primes = append(*primes, i)
		}
	}
}
