package implementation

import "sort"

func cut_the_sticks(length []int) []int {
	sort.Ints(length)
	ret := make([]int, 0)
	ret = append(ret, len(length))
	cnt := 0
	for i := 0; i < len(length); i++ {
		if i == 0 {
			cnt = 1
		} else if i == len(length)-1 {
			if length[i] != length[i-1] {
				ret = append(ret, 1)
			} else {
				ret = append(ret, cnt+1)
			}

		} else if length[i] != length[i-1] {
			ret = append(ret, cnt)
			cnt = 1
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
