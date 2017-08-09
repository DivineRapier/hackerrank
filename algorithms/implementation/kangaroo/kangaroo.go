package kangaroo

func kangaroo(x1, v1, x2, v2 int) bool {
	if v1 == v2 {
		return x1 == x2
	}
	return ((x1-x2)%(v1-v2) == 0) && ((x1-x2)/(v2-v1) > 0)
}
