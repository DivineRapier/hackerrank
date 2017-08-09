package sock_merchant

import "testing"

func TestSockMerchant(t *testing.T) {
	if a := sock_merchant([]int{10, 20, 20, 10, 10, 30, 50, 10, 20}); a != 3 {
		t.Errorf("%d\n", a)
	}
}
