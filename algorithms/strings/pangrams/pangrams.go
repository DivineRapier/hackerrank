package pangrams

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.hackerrank.com/challenges/pangrams

func pangrams(s []byte) string {

	var set [26]bool

	if len(s) < 26 {
		return "not pangram"
	}

	mask, cnt := 1<<5, 0

	for _, v := range s {

		if v == ' ' {
			continue
		}
		// fmt.Println(i, v)
		if !set[int(v)|mask-97] {
			set[int(v)|mask-97] = true
			cnt++
		}
	}

	if cnt == 26 {
		return "pangram"
	}
	return "not pangram"
}

func main() {

	var reader = bufio.NewReader(os.Stdin)

	bs, _, _ := reader.ReadLine()

	// fmt.Printf("%s  %v   %v", bs, a, b)

	// bs := []byte("We promptly judged antique ivory buckles for the next prize ")
	fmt.Println(pangrams(bs))
}
