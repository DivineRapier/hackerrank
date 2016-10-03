package warmup

import "testing"

func TestMaxPerimeterTriangle(t *testing.T) {
	if a, b, c := maximum_perimeter_triangle([]int{1, 1, 1, 3, 3}); a != 1 || b != 3 || c != 3 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
}
