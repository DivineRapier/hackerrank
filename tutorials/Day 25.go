package Tutorials

//import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	} else if num == 2 {
		return true
	} else if num%2 == 0 {
		return true
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// func main() {
//  //Enter your code here. Read input from STDIN. Print output to STDOUT
//     var T, N int
//     fmt.Scan(&N)

//     for i:=0; i<T; i++ {
//         fmt.Scan(&N)
//         if isPrime(N) {
//             fmt.Println("Prime")
//         } else {
//             fmt.Println("Not Prime")
//         }
//     }
// }
