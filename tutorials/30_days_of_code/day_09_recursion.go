package tutorials

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
    
	return n * factorial(n-1)
}


func recursion() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    
   var n int
	fmt.Scan(&n)

	fmt.Print(factorial(n))
}