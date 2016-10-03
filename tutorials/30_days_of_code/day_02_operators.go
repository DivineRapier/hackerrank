package tutorials

import "fmt"

func operators() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var (
		cost float64
		tip  int
		tax  int
	)

	fmt.Scanf("%f", &cost)
	fmt.Scanf("%d", &tip)
	fmt.Scanf("%d", &tax)

	total := int(cost + cost*float64(tip+tax)/100 + 0.5)

	fmt.Printf("The total meal cost is %d dollars.", total)

}
