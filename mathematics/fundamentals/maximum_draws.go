package fundamentals

import "fmt"

func maximum_draws() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var cases, pairs int
	fmt.Scanf("%d", &cases)

	for i := 0; i < cases; i++ {
		fmt.Scanf("%d", &pairs)
		//fmt.Printf("----\n")
		fmt.Printf("%d\n", pairs+1)
	}

}
