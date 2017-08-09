package cut_the_sticks

import "sort"

func cut_the_sticks(length []int) []int {
	sort.Ints(length)
	ret := make([]int, 0)
	remain_cnt := len(length)
	ret = append(ret, remain_cnt)
	cnt := 0
	for i := 0; i < len(length) && remain_cnt > 0; i++ {
		// 如果是第一个 数量直接设置为1
		switch {
		case i == 0:
			cnt = 1
		case i == len(length)-1:
			if length[i] != length[i-1] {
				ret = append(ret, 1)
			}
		case length[i] != length[i-1]:
			remain_cnt -= cnt
			ret = append(ret, remain_cnt)
			cnt = 1
		default:
			cnt++
		}
	}
	return ret
}

func same_slice(sli1, sli2 []int) bool {
	if len(sli1) != len(sli2) {
		return false
	}
	for i := 0; i < len(sli1); i++ {
		if sli1[i] != sli2[i] {
			return false
		}
	}
	return true
}
