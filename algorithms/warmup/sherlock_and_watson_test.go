package warmup

import "testing"

func TestSherlockAndWatson(t *testing.T) {
	if a := sherlock_and_watson(3, 2, 0); a != 1 {
		t.Errorf("final at %d\n", a)
	}
}
