package tutorials

import "testing"

func EXPECT_EQ(n, k, max int, t *testing.T) {
	if a := getMax(n, k); a != max {
		t.Errorf("期望的最大值是: %d, 函数执行结果为: %d\n", max, getMax(n, k))
	}
}

func TestBitwiseAnd(t *testing.T) {
	EXPECT_EQ(5, 2, 1, t)
	EXPECT_EQ(8, 5, 4, t)
	EXPECT_EQ(2, 2, 0, t)
	EXPECT_EQ(5, 2, 1, t)
	EXPECT_EQ(5, 2, 1, t)

}
