package extra_long_factorials

import "testing"

func TestExtraLongFactorials(t *testing.T) {
	var n int64
	n = 1
	if a := extra_long_factorials(n); a != "1" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 2
	if a := extra_long_factorials(n); a != "2" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 3
	if a := extra_long_factorials(n); a != "6" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 4
	if a := extra_long_factorials(n); a != "24" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 5
	if a := extra_long_factorials(n); a != "120" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 1
	if a := extra_long_factorials(n); a != "1" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 1
	if a := extra_long_factorials(n); a != "1" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 1
	if a := extra_long_factorials(n); a != "1" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 1
	if a := extra_long_factorials(n); a != "1" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
	n = 1
	if a := extra_long_factorials(n); a != "1" {
		t.Errorf("a: %d\tfactorials: %s\n", n, a)
	}
}
