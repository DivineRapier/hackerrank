package fundamentals

import "testing"

func TestHandshake(t *testing.T) {
	if a := handshake(1); a != 0 {
		t.Errorf("%d expected, but %d", 0, a)
	}
	if a := handshake(2); a != 1 {
		t.Errorf("%d expected, but %d", 1, a)
	}
	if a := handshake(2924); a != 4273426 {
		t.Errorf("%d expected, but %d", 4273426, a)
	}

	if a := handshake(7062); a != 24932391 {
		t.Errorf("%d expected, but %d", 25236960, a)
	}

	if a := handshake(7105); a != 25236960 {
		t.Errorf("%d expected, but %d", 25236960, a)
	}

	if a := handshake(4716); a != 11117970 {
		t.Errorf("%d expected, but %d", 11117970, a)
	}

	if a := handshake(9111); a != 41500605 {
		t.Errorf("%d expected, but %d", 41500605, a)
	}

	if a := handshake(6913); a != 23891328 {
		t.Errorf("%d expected, but %d", 23891328, a)
	}

	if a := handshake(1421); a != 1008910 {
		t.Errorf("%d expected, but %d", 1008910, a)
	}

	if a := handshake(9429); a != 44448306 {
		t.Errorf("%d expected, but %d", 44448306, a)
	}

	if a := handshake(2378); a != 2826253 {
		t.Errorf("%d expected, but %d", 2826253, a)
	}

	if a := handshake(4540); a != 10303530 {
		t.Errorf("%d expected, but %d", 10303530, a)
	}
}
