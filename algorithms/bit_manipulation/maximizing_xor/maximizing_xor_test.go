package bit_manipulation

import "testing"

func TestMaximizingXor(t *testing.T) {
	if a := maximizing_xor(1, 10); a != 15 {
		t.Errorf("expect %d, but %d\n", 15, a)
	}
	if a := maximizing_xor(5, 6); a != 3 {
		t.Errorf("expect %d, but %d\n", 3, a)
	}
	if a := maximizing_xor(10, 15); a != 7 {
		t.Errorf("expect %d, but %d\n", 7, a)
	}
}
