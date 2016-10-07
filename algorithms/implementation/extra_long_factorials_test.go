package implementation

import "testing"

func TestExtraLongFactorials(t *testing.T) {
	var n int64
	n = 1
	if a := extra_long_factorials(n); a.Int64() != 1 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 2
	if a := extra_long_factorials(n); a.Int64() != 2 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 3
	if a := extra_long_factorials(n); a.Int64() != 6 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 4
	if a := extra_long_factorials(n); a.Int64() != 24 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 5
	if a := extra_long_factorials(n); a.Int64() != 120 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 1
	if a := extra_long_factorials(n); a.Int64() != 1 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 1
	if a := extra_long_factorials(n); a.Int64() != 1 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 1
	if a := extra_long_factorials(n); a.Int64() != 1 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 1
	if a := extra_long_factorials(n); a.Int64() != 1 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
	n = 1
	if a := extra_long_factorials(n); a.Int64() != 1 {
		t.Errorf("a: %d\tfactorials: %s\n", n, a.String())
	}
}
