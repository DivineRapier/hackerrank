package arrays_left_rotation

/*
5 4
[1, 2, 3, 4, 5] -> [5, 1, 2, 3, 4]
*/

func arrays_left_rotation(arr []int, rotation int) []int {
	ret := make([]int, len(arr), len(arr))
	index := 0
	for i := rotation; i < len(arr); i++ {
		ret[index] = arr[i]
		index++
	}
	for i := 0; i < rotation; i++ {
		ret[index] = arr[i]
		index++
	}
	return ret
}

func same_slice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
