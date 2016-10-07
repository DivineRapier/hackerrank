package sort

import "testing"

func TestRunningTimeOfAlgorithms(t *testing.T) {
	if a := running_time_of_algorithms([]int{2, 1, 3, 1, 2}); a != 4 {
		t.Errorf("%d\n", a)
	}
}
