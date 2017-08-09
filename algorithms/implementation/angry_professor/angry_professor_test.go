package angry_professor

import "testing"

func TestAngryProfessor(t *testing.T) {
	a := []int{-1, -3, 4, 2}
	b := 3
	if !angry_professor(a, b) {
		t.Errorf("%v %v should be YES\n", a, b)
	}
	a = []int{0, -1, 2, 1}
	b = 2
	if angry_professor(a, b) {
		t.Errorf("%v %v should be NO\n", a, b)
	}
}
