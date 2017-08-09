package reverse_game

import "testing"

func TestReverseGame(t *testing.T) {
	if a := reverse_game(3, 1); a != 2 {
		t.Errorf("%d\n", a)
	}
	if a := reverse_game(5, 2); a != 4 {
		t.Errorf("%d\n", a)
	}
}
