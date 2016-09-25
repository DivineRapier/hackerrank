package fundamentals

/*


1   3   5   7   9   0
0   2   4   6   8   1
*/

func strange_grid_again(row, col int) int {
	n := (row - 1) >> 1
	dec := n * 10
	dig := 0
	if row&0x1 == 0 {
		dig = 2*col - 1
	} else {
		dig = 2 * (col - 1)
	}
	return dec + dig
}
