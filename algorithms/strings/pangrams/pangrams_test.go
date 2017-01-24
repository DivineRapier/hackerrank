package strings

import "testing"

func TestPangrams(t *testing.T) {
	if pangrams("We promptly judged antique ivory buckles for the next prize ") != "pangram" {
		t.Errorf("error")
	}
	if pangrams("We promptly judged antique ivory buckles for the prize ") == "pangram" {
		t.Errorf("error")
	}
	if pangrams("We promptly judged antique ivory buckles for the next prize") != "pangram" {
		t.Errorf("error")
	}

}
