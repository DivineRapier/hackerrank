package cracking_the_coding_interview

import "testing"

func TestSortingBubbleSort(t *testing.T) {
	if ret := sorting_bubble_sort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}); !same_slice(ret, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("%v\n", ret)
	}
}

func TestSortingBubbleSortCount(t *testing.T) {
	if a, b := sorting_bubble_sort_count([]int{4, 3, 2, 1, 0}); !same_slice(b, []int{0, 1, 2, 3, 4}) || a != 10 {
		t.Errorf("%v\t%v\n", a, b)
	}
	if a, b := sorting_bubble_sort_count([]int{1, 2, 3}); !same_slice(b, []int{1, 2, 3}) || a != 0 {
		t.Errorf("%v\t%v\n", a, b)
	}
	if a, b := sorting_bubble_sort_count([]int{3, 2, 1}); !same_slice(b, []int{1, 2, 3}) || a != 3 {
		t.Errorf("%v\t%v\n", a, b)
	}
}
