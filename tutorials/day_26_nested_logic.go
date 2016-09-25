package tutorials

import "fmt"

// Date ...
type Date struct {
	Day   int64
	Month int64
	Year  int64
}

func computeFine(actual, expected *Date) int64 {
	if (actual.Year - expected.Year) < 0 {
		return 0
	}
	if y := actual.Year - expected.Year; y > 0 {
		return 10000
	} else if m := actual.Month - expected.Month; m > 0 {
		return 500 * m
	} else if d := actual.Day - expected.Day; d > 0 {
		return 15 * d
	}
	return 0
}

func nested_logic() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var actual, expected Date
	fmt.Scanf("%d%d%d", &actual.Day, &actual.Month, &actual.Year)
	fmt.Scanf("%d%d%d", &expected.Day, &expected.Month, &expected.Year)

	fmt.Println(computeFine(&actual, &expected))
}
