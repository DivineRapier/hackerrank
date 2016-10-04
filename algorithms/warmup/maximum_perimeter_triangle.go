package warmup

import "sort"

func is_perimeter_triangle(a, b, c int) bool {
	return (a == b && a+b > c) || (a == c && a+c > b) || (b == c && b+c > a)
}

// 返回可以构成周长最大的等边三角形的变长 从小到大
func maximum_perimeter_isosceles_triangle(edges []int) (int, int, int) {
	sort.Ints(edges)
	length := len(edges)
	if edges[length-1] == edges[length-2] {
		return edges[length-3], edges[length-2], edges[length-1]
	}
	for i := length - 2; i >= 1; i-- {
		if edges[i] == edges[i-1] {
			if 2*edges[i] <= edges[i+1] {
				if i >= 2 {
					return edges[i-2], edges[i-1], edges[i]
				} else {
					break
				}
			} else {
				for j := i + 1; j < length; j++ {
					if (edges[i] << 1) > edges[j] {
						if (j+1 == length) || ((edges[i] << 1) <= edges[j+1]) {
							return edges[i], edges[i], edges[j]
						}
					}
				}
			}

		}
	}
	return -1, -1, -1
}

func maximum_perimeter_triangle(edges []int) (int, int, int) {
	sort.Ints(edges)
	length := len(edges)
	if length < 3 {
		return -1, -1, -1
	}
	for ; length >= 3; length-- {
		a, b, c := edges[length-1], edges[length-2], edges[length-3]
		//log.Errorf("a: %d\nb: %d\nc: %d\n", a, b, c)
		if a-b < c {
			return c, b, a
		}
	}
	return -1, -1, -1
}
