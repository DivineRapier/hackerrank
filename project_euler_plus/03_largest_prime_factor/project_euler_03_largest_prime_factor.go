package main

import (
	"fmt"
	"math"
)

var PrimeList = []uint64{2, 3, 5, 7, 11, 13}

func init() {

	for i := uint64(1); i < math.MaxUint16; i++ {
		if IsPrime(i) {
			PrimeList = append(PrimeList, i)
		}
	}

}

func IsPrime(a uint64) bool {
	for _, v := range PrimeList {
		if v >= a {
			return false
		}

		if a%v == 0 {
			return false
		}
	}

	return true
}

func FindLargestPrimeFactor(a uint64) uint64 {

	privot := uint64(math.Sqrt(float64(a)))

	for _, v := range PrimeList {

		if v > privot {
			return a
		}

		if a%v == 0 {
			return a / v
		}
	}
	return a
}

func main() {

	var t, n uint64

	fmt.Scan(&t)

	for i := uint64(0); i < t; i++ {
		fmt.Scan(&n)
		fmt.Println(FindLargestPrimeFactor(n))
	}
}
