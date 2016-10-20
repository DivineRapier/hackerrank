package bit_manipulation

import "testing"

func TestLonelyInteger(t *testing.T) {
	if a := lonely_integer([]int{0, 0, 1, 2, 1}); a != 2 {
		t.Errorf("expect %d but %d\n", 2, a)
	}
	if a := lonely_integer([]int{1, 2, 1}); a != 2 {
		t.Errorf("expect %d but %d\n", 2, a)
	}
	if a := lonely_integer([]int{1}); a != 1 {
		t.Errorf("expect %d but %d\n", 1, a)
	}
	if a := lonely_integer([]int{4, 9, 95, 93, 57, 4, 57, 93, 9}); a != 95 {
		t.Errorf("expect %d but %d\n", 1, a)
	}
}
