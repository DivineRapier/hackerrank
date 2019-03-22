package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/recursive-digit-sum/problem

// Complete the superDigit function below.
func superDigit(n string, k int32) int32 {
	return int32(sumNInt64(sumNString(n) * int64(k)))
}

func sumNString(n string) int64 {
	var result int64
	length := len(n)
	for i := 0; i < length; i++ {
		result += int64(n[i] - '0')
	}
	return result
}

func sumNInt64(n int64) int64 {
	fmt.Println("n:", n)
	if n < 10 {
		return n
	}
	var result int64
	for n > 0 {
		result += (n % 10)
		n /= 10
	}
	return sumNInt64(result)
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	n := nk[0]

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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
