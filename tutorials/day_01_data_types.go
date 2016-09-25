package tutorials

import (
	"bufio"
	"fmt"
	"os"
)

func DataType() {
	var i = 4
	var d = 4.0
	var s = "HackerRank "

	scanner := bufio.NewReader(os.Stdin)

	// Declare second integer, double, and String variables.
	var (
		a   int
		b   float64
		c   string
		err error
	)

	// Read and save an integer, double, and String to your variables.
	fmt.Scanf("%d", &a)
	fmt.Scanf("%f", &b)
	c, err = scanner.ReadString('\n')

	if err != nil {

	}

	// Print the sum of both integer variables on a new line.
	fmt.Printf("%d\n", i+a)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+b)
	// Concatenate and print the String variables on a new line
	fmt.Printf("%s%s\n", s, c)
	// The 's' variable above should be printed first.
}
