package arrays_left_rotation

import "testing"

func TestLeftRotation(t *testing.T) {
	if ret := arrays_left_rotation([]int{1, 2, 3, 4, 5}, 4); !same_slice(ret, []int{5, 1, 2, 3, 4}) {
		t.Errorf("error: %v\n", ret)
	}
}
