package super_reduced_string

import "testing"

func TestSuperReducedString(t *testing.T) {
	if r := super_reduced_string("aaabccddd"); r != "abd" {
		t.Errorf("error: unexpected abd, but %s\n", r)
	}
	if r := super_reduced_string("abba"); r != "" {
		t.Errorf("error: unexpected empty, but %s\n", r)
	}
	if r := super_reduced_string(""); r != "" {
		t.Errorf("error: unexpected empty, but %s\n", r)
	}
	if r := super_reduced_string("aa"); r != "" {
		t.Errorf("error: unexpected empty, but %s\n", r)
	}
}
