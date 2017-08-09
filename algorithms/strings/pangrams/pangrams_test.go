package pangrams

import "testing"

func TestPangrams(t *testing.T) {
	if pangrams([]byte("We promptly judged antique ivory buckles for the next prize ")) != "pangram" {
		t.Errorf("error")
	}
	if pangrams([]byte("We promptly judged antique ivory buckles for the prize ")) != "not pangram" {
		t.Errorf("error")
	}
	if pangrams([]byte("We promptly judged antique ivory buckles for the next prize")) != "pangram" {
		t.Errorf("error")
	}
	if pangrams([]byte("We promptly judged antique ivory buckles for the next prize")) != "pangram" {
		t.Errorf("error")
	}

}
