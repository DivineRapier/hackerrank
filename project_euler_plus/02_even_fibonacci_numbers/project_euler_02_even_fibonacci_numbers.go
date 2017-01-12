package main

import (
	"fmt"
	"math"
)

var (
	// FiboSeq a slice to store even fibo numbers
	FiboSeq  []uint64
	SumTable []Node
)

type Node struct {
	LastValue uint64
	SumValue  uint64
}

func init() {
	a, b := uint64(1), uint64(1)

	for b < math.MaxUint64 {
		if a&0x01 == 0 {
			FiboSeq = append(FiboSeq, a)
		}
		if a > math.MaxUint64>>1 {
			break
		}
		a, b = a+b, a
	}

	SumTable = make([]Node, len(FiboSeq))

	var firstN Node
	firstN.LastValue = FiboSeq[0]
	firstN.SumValue = FiboSeq[0]
	SumTable[0] = firstN

	for i := 1; i < len(FiboSeq); i++ {

		SumTable[i].LastValue = FiboSeq[i]
		SumTable[i].SumValue = FiboSeq[i] + SumTable[i-1].SumValue
	}

}

func search(n uint64) uint64 {
	if n <= 1 {
		return 0
	}
	for i := 0; i < len(SumTable); i++ {
		if SumTable[i].LastValue > n {
			return SumTable[i-1].SumValue
		}
	}
	return SumTable[len(SumTable)-1].SumValue
}

func main() {
	fmt.Println(FiboSeq)
	fmt.Println()
	fmt.Println(SumTable)
}
