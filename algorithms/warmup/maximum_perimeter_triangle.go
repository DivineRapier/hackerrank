package warmup

import "sort"

func is_perimeter_triangle(a, b, c int) bool {
	return (a == b && a+b > c) || (a == c && a+c > b) || (b == c && b+c > a)
}

func maximum_perimeter_triangle(edges []int) (int, int, int) {
	sort.Ints(edges)
	length := len(edges)
	for a, b, c := edges[length-1], edges[length-2], edges[length-3]; c >= 0; a, b, c = a-1, b-1, c-1 {
		if is_perimeter_triangle(a, b, c) {
			return c, b, a
		}
	}
	return -1, -1, -1
}
