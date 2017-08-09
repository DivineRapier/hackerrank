package sum_vs_xor

import "testing"

// func TestSumVSXor(t *testing.T) {
// 	if a := sum_vs_xor(5); a != 2 {
// 		t.Errorf("expect %d but %d\n", 2, a)
// 	}
// 	if a := sum_vs_xor(10); a != 4 {
// 		t.Errorf("expect %d but %d\n", 4, a)
// 	}
// }

func TestZeroCount(t *testing.T) {
	if a := zero_count(0); a != 0 {
		t.Errorf("expect %d but %d\n", 0, a)
	}
	if a := zero_count(1); a != 0 {
		t.Errorf("expect %d but %d\n", 0, a)
	}
	if a := zero_count(2); a != 1 {
		t.Errorf("expect %d but %d\n", 1, a)
	}
	if a := zero_count(3); a != 0 {
		t.Errorf("expect %d but %d\n", 0, a)
	}
	if a := zero_count(4); a != 2 {
		t.Errorf("expect %d but %d\n", 2, a)
	}
	if a := zero_count(5); a != 1 {
		t.Errorf("expect %d but %d\n", 1, a)
	}
}

func TestSumXor(t *testing.T) {
	if a := sum_vs_xor(0); a != 1 {
		t.Errorf("expect %d but %d\n", 1, a)
	}
	if a := sum_vs_xor(1); a != 1 {
		t.Errorf("expect %d but %d\n", 4, a)
	}
	if a := sum_vs_xor(2); a != 2 {
		t.Errorf("expect %d but %d\n", 4, a)
	}
	if a := sum_vs_xor(3); a != 1 {
		t.Errorf("expect %d but %d\n", 4, a)
	}
	if a := sum_vs_xor(4); a != 4 {
		t.Errorf("expect %d but %d\n", 4, a)
	}
	if a := sum_vs_xor(5); a != 2 {
		t.Errorf("expect %d but %d\n", 4, a)
	}
	if a := sum_vs_xor(6); a != 2 {
		t.Errorf("expect %d but %d\n", 2, a)
	}
	if a := sum_vs_xor(7); a != 1 {
		t.Errorf("expect %d but %d\n", 1, a)
	}
	if a := sum_vs_xor(8); a != 8 {
		t.Errorf("expect %d but %d\n", 8, a)
	}
	if a := sum_vs_xor(9); a != 4 {
		t.Errorf("expect %d but %d\n", 4, a)
	}
	if a := sum_vs_xor(10); a != 4 {
		t.Errorf("expect %d but %d\n", 4, a)
	}
}
