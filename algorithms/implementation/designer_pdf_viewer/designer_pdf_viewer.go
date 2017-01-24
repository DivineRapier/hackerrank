package main

func contiguous_block_area(height []int, chars []byte) int {

	if len(height) != 26 {
		panic("increct width")
	}

	max := 0

	for _, v := range chars {
		if max < height[int(v-'a')] {
			max = height[int(v-'a')]
		}
	}

	return len(chars) * max
}
