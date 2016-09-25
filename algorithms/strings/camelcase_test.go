package strings

import "testing"

func TestCamelcase(t *testing.T) {
	if n := camelcase("saveChangesInTheEditor"); n != 5 {
		t.Errorf("5 words expected, but %d\n", n)
	}
}
