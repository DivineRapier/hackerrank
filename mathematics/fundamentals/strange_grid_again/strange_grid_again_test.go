package strange_grid_again

import "testing"

func TestStrangeGridAgain(t *testing.T) {
	if a := strange_grid_again(6, 3); a != 25 {
		t.Errorf("wrong! %d\n", a)
	}
}
