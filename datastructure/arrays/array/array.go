package main

import (
	"fmt"
)

// DS ...
// An array is a type of data structure that stores elements of the same type in a contiguous block of memory. In an array, , of size , each memory location has some unique index,  (where ), that can be referenced as  (you may also see it written as ).
//
// Given an array, , of  integers, print each element in reverse order as a single line of space-separated integers.
//
// Note: If you've already solved our C++ domain's Arrays Introduction challenge, you may want to skip this.
//
// Input Format
//
// The first line contains an integer,  (the number of integers in ).
// The second line contains  space-separated integers describing .
//
// Constraints
//
// Output Format
//
// Print all  integers in  in reverse order as a single line of space-separated integers.
//
// Sample Input
//
// 4
// 1 4 3 2
// Sample Output
//
// 2 3 4 1
func DS() {
	var count int
	fmt.Scanf("%d", &count)
	array := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &array[i])
	}

	for i := count - 1; i >= 0; i-- {
		fmt.Printf("%d ", array[i])
	}

}

// DS2D ...
// Context
// Given a  2D Array, :
//
// 1 1 1 0 0 0
// 0 1 0 0 0 0
// 1 1 1 0 0 0
// 0 0 0 0 0 0
// 0 0 0 0 0 0
// 0 0 0 0 0 0
// We define an hourglass in  to be a subset of values with indices falling in this pattern in 's graphical representation:
//
// a b c
//   d
// e f g
// There are  hourglasses in , and an hourglass sum is the sum of an hourglass' values.
//
// Task
// Calculate the hourglass sum for every hourglass in , then print the maximum hourglass sum.
//
// Note: If you have already solved the Java domain's Java 2D Array challenge, you may wish to skip this challenge.
//
// Input Format
//
// There are  lines of input, where each line contains  space-separated integers describing 2D Array ; every value in  will be in the inclusive range of  to .
//
// Constraints
//
// Output Format
//
// Print the largest (maximum) hourglass sum found in .
//
// Sample Input
//
// 1 1 1 0 0 0
// 0 1 0 0 0 0
// 1 1 1 0 0 0
// 0 0 2 4 4 0
// 0 0 0 2 0 0
// 0 0 1 2 4 0
// Sample Output
//
// 19
// Explanation
//
//  contains the following hourglasses:
//
// 1 1 1   1 1 0   1 0 0   0 0 0
//   1       0       0       0
// 1 1 1   1 1 0   1 0 0   0 0 0
//
// 0 1 0   1 0 0   0 0 0   0 0 0
//   1       1       0       0
// 0 0 2   0 2 4   2 4 4   4 4 0
//
// 1 1 1   1 1 0   1 0 0   0 0 0
//   0       2       4       4
// 0 0 0   0 0 2   0 2 0   2 0 0
//
// 0 0 2   0 2 4   2 4 4   4 4 0
//   0       0       2       0
// 0 0 1   0 1 2   1 2 4   2 4 0
// The hourglass with the maximum sum () is:
//
// 2 4 4
//   2
// 1 2 4
func DS2D() {
	var sum, maxSum int

	array2D := make([][]int, 6)

	for i := 0; i < 6; i++ {
		array2D[i] = make([]int, 6)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scanf("%d", &array2D[i][j])
		}
	}

	for i := 1; i < 6-1; i++ {
		for j := 1; j < 6-1; j++ {
			sum = array2D[i][j] + array2D[i-1][j-1] + array2D[i-1][j] + array2D[i-1][j+1] + array2D[i+1][j-1] + array2D[i+1][j] + array2D[i+1][j+1]
			if maxSum < sum {
				maxSum = sum
			}
		}
	}
	fmt.Println(maxSum)
}
