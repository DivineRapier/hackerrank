package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/the-power-sum/problem

// Complete the powerSum function below.
func powerSum(X int32, N int32) int32 {
	return foo(X, N, 1)
}

func foo(X int32, N int32, num int32) int32 {
	value := int32(math.Pow(float64(num), float64(N)))
	if value == X {
		return 1
	}
	if value < X {
		return foo(X, N, num+1) + foo(X-value, N, num+1)
	}
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	XTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	X := int32(XTemp)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	result := powerSum(X, N)

	fmt.Fprintf(writer, "%d\n", result)

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
