package main

/*
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
1, 3, 5, 7, 9, 8, 6, 4, 2, 0        0
-------------------------------
1, 3, 5, 7, 9, 8, 6, 4, 2, 2        0
1, 3, 5, 7, 9, 8, 6, 4, 4, 2        0
1, 3, 5, 7, 9, 8, 6, 6, 4, 2        0
1, 3, 5, 7, 9, 8, 8, 6, 4, 2        0
1, 3, 5, 7, 9, 9, 8, 6, 4, 2        0
1, 3, 5, 7, 7, 9, 8, 6, 4, 2        0
1, 3, 5, 5, 7, 9, 8, 6, 4, 2        0
1, 3, 3, 5, 7, 9, 8, 6, 4, 2        0
0, 1, 3, 5, 7, 9, 8, 6, 4, 2        0

0, 1, 3, 5, 7, 9, 8, 6, 4, 4        0
0, 1, 3, 5, 7, 9, 8, 6, 6, 4        0
0, 1, 3, 5, 7, 9, 8, 8, 6, 4        0
0, 1, 3, 5, 7, 9, 9, 8, 6, 4        0
0, 1, 3, 5, 7, 7, 9, 8, 6, 4        0
0, 1, 3, 5, 5, 7, 9, 8, 6, 4        0
0, 1, 3, 3, 5, 7, 9, 8, 6, 4        0
0, 1, 2, 3, 5, 7, 9, 8, 6, 4        0

*/

func insertion_sort_partI(nums []int) [][]int {
	ret := make([][]int, 0)
	length := len(nums)
	for i := 0; i < length; i++ {
		tar := nums[length-1]
		for j := length - 2; j >= i; j-- {

			if nums[j] > tar {
				nums[j+1] = nums[j]
				if j == i {
					tmp := make([]int, length)
					copy(tmp, nums)
					if len(ret) == 0 || !same_slice(ret[len(ret)-1], tmp) {
						ret = append(ret, tmp)
					}
					nums[i] = tar
				}
			} else {
				nums[j+1] = tar
				tar = nums[j]
			}

			tmp := make([]int, length)
			copy(tmp, nums)
			if len(ret) == 0 || !same_slice(ret[len(ret)-1], tmp) {
				ret = append(ret, tmp)
			}
		}

		//fmt.Println(nums)
	}
	return ret
}

func is_sorted(arr []int) bool {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			//fmt.Println(arr)
			return false
		}
	}
	return true
}

func same_slice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
