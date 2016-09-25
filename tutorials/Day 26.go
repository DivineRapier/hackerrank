package Tutorials

// Date ...
type Date struct {
	Day   int64
	Mouth int64
	Year  int64
}

func computeFine(actual, expected *Date) int64 {
	if y := actual.Year - expected.Year; y > 0 {
		return 10000
	}
	if m := actual.Mouth - expected.Mouth; m > 0 {
		return 500 * m
	}
	if d := actual.Day - expected.Day; d > 0 {
		return 15 * d
	}
	return 0
}
