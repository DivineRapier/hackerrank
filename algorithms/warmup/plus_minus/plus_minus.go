package plus_minus

func plus_minus(arr []int) (float64, float64, float64) {
	cntpos, cntneg, cntzer := 0, 0, 0
	l := len(arr)
	for _, v := range arr {
		if v < 0 {
			cntneg++
		} else if v == 0 {
			cntzer++
		} else {
			cntpos++
		}
	}
	return float64(cntpos) / float64(l), float64(cntneg) / float64(l), float64(cntzer) / float64(l)
}
