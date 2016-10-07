package sort

import "testing"

func TestIntroToTutorialChallenges(t *testing.T) {
	if a := intro_to_tutorial_challenges([]int{1, 4, 5, 7, 9, 12}, 4); a != 1 {
		t.Errorf("a = %d\n", a)
	}
}
