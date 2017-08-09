package divisible_sum_pairs

import "testing"

func TestDivisibelSumPairs(t *testing.T) {
	if a := divisible_sum_pairs([]int{1, 3, 2, 6, 1, 2}, 3); a != 5 {
		t.Errorf("expect 5 but %d\n", a)
	}
}
