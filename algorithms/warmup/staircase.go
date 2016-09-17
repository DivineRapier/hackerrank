package warmup

import "fmt"

func staircase(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < n-i+1; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
