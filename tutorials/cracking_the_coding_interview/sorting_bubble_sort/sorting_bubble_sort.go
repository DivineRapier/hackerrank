package sorting_bubble_sort

func sorting_bubble_sort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func sorting_bubble_sort_count(arr []int) (int, []int) {
	totalNumber := 0
	for i := 0; i < len(arr); i++ {
		numberOfSwaps := 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				numberOfSwaps++
			}
		}
		totalNumber += numberOfSwaps
		if numberOfSwaps == 0 {
			break
		}
	}
	return totalNumber, arr
}
