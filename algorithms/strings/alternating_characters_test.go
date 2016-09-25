package strings

import "testing"

func TestAlternativeCharaters(t *testing.T) {
	if a := alternating_characters("a"); a != 0 {
		t.Errorf("%d\n", a)
	}
	if a := alternating_characters("aaaa"); a != 3 {
		t.Errorf("%d\n", a)
	}
	if a := alternating_characters("bbbbb"); a != 4 {
		t.Errorf("%d\n", a)
	}
	if a := alternating_characters("abababab"); a != 0 {
		t.Errorf("%d\n", a)
	}
	if a := alternating_characters("bababa"); a != 0 {
		t.Errorf("%d\n", a)
	}
	if a := alternating_characters("aaabbb"); a != 4 {
		t.Errorf("%d\n", a)
	}
}
