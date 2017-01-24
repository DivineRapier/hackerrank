package main

// https://www.hackerrank.com/challenges/circular-array-rotation

func sherlock_and_watson(n, k, tar int) int {
	return ((tar-k)%n + n) % n
}
