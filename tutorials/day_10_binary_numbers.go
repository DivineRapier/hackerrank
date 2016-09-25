package tutorials

import "fmt"

func count(num int) int {
    i := 1
    cur := 0
    max := 0
    for num > (i/2) {
        if (num & i) > 0 {
            cur++
        } else {
            if cur > max {
                max = cur
            }
            cur = 0
        }
        i = i << 1
    }
    return max
}


func binary_numbers() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var num int
    fmt.Scan(&num)
    fmt.Printf("%d\n", count(num))
}
