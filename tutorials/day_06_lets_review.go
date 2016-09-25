package tutorials

import "fmt"
import "bytes"

// divideString ...
func divideString(src string) ([]byte, []byte) {
	var eve, odd bytes.Buffer

	for i := 0; i < len(src); i++ {
		if i&0x1 == 0 {
			eve.WriteByte(src[i])
		} else {
			odd.WriteByte(src[i])
		}
	}
    return eve.Bytes(), odd.Bytes()
}


func review() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s", &s)

      eve, odd := divideString(s)
		fmt.Printf("%s %s\n", eve, odd)
	} 
}