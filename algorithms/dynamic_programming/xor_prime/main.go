package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	primeSet []int
	dataSet  map[int]struct{}
)

func init() {
	primeSet = make([]int, 0, 4000)

	primeSet = append(primeSet, 2, 3, 5, 7, 11, 13, 17, 19)

	for i := 21; i <= 4500; i += 2 {
		prime := true
		for j := 0; j < len(primeSet) && primeSet[j]*primeSet[j] <= i; j++ {
			if i%primeSet[j] == 0 {
				prime = false
				break
			}
		}
		if prime {
			primeSet = append(primeSet, i)
		}
	}
	for i, v := range primeSet {
		if v >= 3500 {
			primeSet = primeSet[i:]
			break
		}
	}
}

func isPrime(x int) bool {
	index := sort.SearchInts(primeSet, x)
	return primeSet[index] == x
}

func main() {
	var (
		testCnt int
		dataCnt int
	)

	inbuf := bufio.NewReader(os.Stdin)

	line, _, _ := inbuf.ReadLine()
	testCnt, _ = strconv.Atoi(string(line))

	for i := 0; i < testCnt; i++ {
		result := 0
		dataSet = make(map[int]struct{})
		line, _, _ = inbuf.ReadLine()
		dataCnt, _ = strconv.Atoi(string(line))
		_ = dataCnt
		line, _, _ = inbuf.ReadLine()
		pieces := bytes.Split(line, []byte(" "))
		for _, piece := range pieces {
			if len(piece) == 0 {
				continue
			}
			num, _ := strconv.Atoi(string(piece))
			if isPrime(num) {
				dataSet[num] = struct{}{}
			}
			result = result ^ num
		}
		if isPrime(result) {
			fmt.Println(len(dataSet) + 1)
		} else {
			fmt.Println(len(dataSet))
		}
	}
}
