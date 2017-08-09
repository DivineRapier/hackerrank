package insertion_sort_partII

func insertion_sort_partII(arr []int) [][]int {
	ret := make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		ret = append(ret, tmp)
	}
	return ret
}
