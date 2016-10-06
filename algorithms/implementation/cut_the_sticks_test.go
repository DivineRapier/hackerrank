package implementation

import "testing"

func TestCutTheSticks(t *testing.T) {
	if a := cut_the_sticks([]int{1, 2, 3, 4, 3, 3, 2, 1}); !same_slice(a, []int{8, 6, 4, 1}) {
		t.Errorf("%v\n", a)
	}
	if a := cut_the_sticks([]int{1, 2, 3, 4, 3, 3, 4, 2, 1}); !same_slice(a, []int{9, 7, 5, 2}) {
		t.Errorf("%v\n", a)
	}
}
