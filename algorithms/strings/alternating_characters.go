package strings

func alternating_characters(str string) int {
	cnt := 0
	l := len(str)
	if l <= 1 {
		return 0
	}
	for i := 0; i < l-1; i++ {
		if str[i] == str[i+1] {
			cnt++
		}
	}
	return cnt
}
