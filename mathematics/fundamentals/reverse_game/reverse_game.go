package reverse_game

// https://www.hackerrank.com/challenges/reverse-game

func reverse_game(n, target int) int {
	if target < ((n - 1) >> 1) {
		return 2*target + 1
	}
	return 2 * (n - 1 - target)

}
