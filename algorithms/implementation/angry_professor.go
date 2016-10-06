package implementation

// https://www.hackerrank.com/challenges/angry-professor

func angry_professor(times []int, num int) bool {
	ok := true
	cnt := 0
	for i := 0; i < len(times); i++ {
		if times[i] <= 0 {
			cnt++
			if cnt == num {
				ok = false
				break
			}
		}
	}
	return ok
}
