package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	length := int64(len(s))
	times := n / length
	mod := n % length
	count1, count2 := countA(s, mod)
	return count1 + count2*times
}

func countA(s string, mod int64) (int64, int64) {
	count1, count2 := int64(0), int64(0)
	for i := int64(0); i < mod; i++ {
		if s[i] == 'a' {
			count1++
		}
	}
	for i := int64(mod); i < int64(len(s)); i++ {
		if s[i] == 'a' {
			count2++
		}
	}
	return count1, count1 + count2
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
