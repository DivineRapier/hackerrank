package main

// https://www.hackerrank.com/challenges/compare-the-triplets?h_r=next-challenge&h_v=zen

func compare_the_triplets(alice, bob []int) (int, int) {
	la, lb := len(alice), len(bob)
	if la != lb {
		return 0, 0
	}
	a, b := 0, 0
	for i := 0; i < la; i++ {
		if alice[i] < bob[i] {
			b++
		} else if alice[i] > bob[i] {
			a++
		}
	}
	return a, b
}
