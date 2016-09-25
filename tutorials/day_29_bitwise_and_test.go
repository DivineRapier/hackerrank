package tutorials

import "testing"

func ExpectEq(n, k, max int, t *testing.T) {
	if a := getMax(n, k); a != max {
		t.Errorf("期望的最大值是: %d, 函数执行结果为: %d\n", max, getMax(n, k))
	}
}

func TestBitwiseAnd(t *testing.T) {
	ExpectEq(5, 2, 1, t)
	ExpectEq(8, 5, 4, t)
	ExpectEq(2, 2, 0, t)
	ExpectEq(5, 2, 1, t)
	ExpectEq(5, 2, 1, t)

}
