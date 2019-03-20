package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/equal/copy-from/102512529

// Complete the equal function below.
func equal(arr []int32) int32 {
	min := 0x7fffffff
	for _, num := range arr {
		if num < int32(min) {
			min = int(num)
		}
	}
	result := 0x7fffffff
	for i := 0; i < 4; i++ {
		res := foo(arr, int32(min), int32(i))
		if res < int32(result) {
			result = int(res)
		}
	}
	return int32(result)
}

func foo(arr []int32, min int32, offset int32) int32 {
	var result int32
	for _, num := range arr {
		result += bar(num - min + offset)
	}
	return result
}

func bar(n int32) int32 {
	var result int32
	result += n / 5
	n %= 5
	result += n / 2
	n %= 2
	result += n
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := equal(arr)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
