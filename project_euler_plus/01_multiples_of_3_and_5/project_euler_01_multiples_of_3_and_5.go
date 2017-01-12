package main

/*
Input Format

First line contains  that denotes the number of test cases. This is followed by  lines, each containing an integer, .

Output Format

For each test case, print an integer that denotes the sum of all the multiples of  3 or 5  below N.
*/

func MultiplesOf3Or5(n int) int {
	countFor3 := (n - 1) / 3
	countFor5 := (n - 1) / 5
	countFor15 := (n - 1) / 15
	return SumOfN(3, countFor3) + SumOfN(5, countFor5) - SumOfN(15, countFor15)
}

func SumOfN(d, n int) int {
	return d * n * (n + 1) / 2
}

func main() {

}
