package insertion_sort_partII

import "testing"

func TestInsertionSortII(t *testing.T) {
	if a := insertion_sort_partII([]int{1, 4, 3, 5, 6, 2}); len(a) != 5 {
		t.Errorf("%d\n", a)
	}
}
