package sort

func running_time_of_algorithms(arr []int) int {
	cnt := 0
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				cnt++
			}
		}
	}
	return cnt
}
