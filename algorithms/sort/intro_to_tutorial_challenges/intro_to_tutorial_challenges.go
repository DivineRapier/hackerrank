package intro_to_tutorial_challenges

func intro_to_tutorial_challenges(arr []int, target int) int {
	a, b := 0, len(arr)
	var mid int
	for a < b {
		mid = a + (b-a)>>1
		if arr[mid] < target {
			a = mid + 1
		} else {
			b = mid
		}
	}
	if arr[a] == target {
		return a
	}
	return -1
}
