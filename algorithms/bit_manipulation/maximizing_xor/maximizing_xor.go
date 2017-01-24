package main

func maximizing_xor(left, right int) int {
	i := 1 << 31
	for left&i == right&i {
		i >>= 1
	}
	return (i<<1 - 1)
}
