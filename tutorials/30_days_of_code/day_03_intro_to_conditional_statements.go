package tutorials

import "fmt"

/*

If  is odd, print Weird
If  is even and in the inclusive range of  to , print Not Weird
If  is even and in the inclusive range of  to , print Weird
If  is even and greater than , print Not Weird

*/


func conditional_statements() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    
    var n int
    fmt.Scanf("%d", &n)
    
    if n & 0x1 == 1 {
        fmt.Printf("Weird\n")
        return 
    }
    
    if n >= 2 && n <=5 {
        fmt.Printf("Not Weird\n")
        return
    }
    
    if n >= 6 && n <= 20 {
        fmt.Printf("Weird\n")
        return 
    }
    
    if n >20 {
        fmt.Printf("Not Weird\n")
        return
    }
    
}