package sort

import "testing"

func TestInsertionSortI(t *testing.T) {
	if a := insertion_sort_partI([]int{2, 4, 6, 8, 3}); len(a) != 4 || !same_slice(a[0], []int{2, 4, 6, 8, 8}) {
		t.Errorf("%v\n", a)
	}
	// if a := insertion_sort_partI([]int{1, 3, 5, 7, 9, 8, 6, 4, 2, 0}); !same_slice(a[0], []int{0, 1, 3, 5, 7, 9, 8, 6, 4, 2}) {
	// 	t.Errorf("%v\n", a)
	// }
	if a := insertion_sort_partI([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}); !same_slice(a[0], []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10}) {
		t.Errorf("%v\n", a)
	}
}

func TestIsSorted(t *testing.T) {
	if !is_sorted([]int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("wrong\n")
	}
	if !is_sorted([]int{1, 2, 3, 4, 5, 6, 6}) {
		t.Errorf("wrong\n")
	}
	if is_sorted([]int{1, 2, 3, 4, 5, 6, 5}) {
		t.Errorf("wrong\n")
	}
}
