package tutorials

import "fmt"

func two_D_arrays() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	arr := make([][]int, 6)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 6)
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	max := -64
	cur := 0
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			cur = arr[i][j] + arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
			if cur > max {
				max = cur
			}
		}
	}

	fmt.Print(max)
}
