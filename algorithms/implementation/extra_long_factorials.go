package implementation

import (
	"math/big"
)

func extra_long_factorials(num int64) string {
	n := big.NewInt(num)
	return n.MulRange(1, num).String()
}
