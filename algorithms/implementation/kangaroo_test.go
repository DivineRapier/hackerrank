package implementation

import "testing"

func TestKangaroo(t *testing.T) {
	if kangaroo(43, 2, 70, 2) {
		t.Errorf("exp false\n")
	}
}
