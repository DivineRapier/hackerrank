package main

import "fmt"

//https://www.hackerrank.com/challenges/find-point

func find_point() {
	var testcases, a, b, c, d int
	fmt.Scanf("%d", &testcases)

	for i := 0; i < testcases; i++ {
		fmt.Scanf("%d%d%d%d", &a, &b, &c, &d)
		fmt.Printf("%d %d\n", 2*c-a, 2*d-b)
	}

}
