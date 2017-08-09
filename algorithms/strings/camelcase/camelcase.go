package camelcase

// https://www.hackerrank.com/challenges/camelcase

func camelcase(input string) int {
	cnt := 0
	l := len(input)
	for i := 0; i < l; i++ {
		if input[i] >= 'A' && input[i] <= 'Z' {
			cnt++
		}
	}
	return cnt + 1
}
