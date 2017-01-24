package main

func mini_max_sum(nums []uint64) (uint64, uint64) {

	nums = nums[:5]

	min, max, sum := nums[0], nums[0], uint64(0)

	for _, v := range nums {
		sum += v

		if min > v {
			min = v
		} else if max < v {
			max = v
		}
	}

	return sum - max, sum - min
}
