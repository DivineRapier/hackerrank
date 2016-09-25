package strings

// https://www.hackerrank.com/challenges/reduced-string

// aaabccdd ->  abd
// baab     ->  bb  ->   ""
// aa       ->  ""

func super_reduced_string(input string) string {

	length := len(input)
	str := make([]byte, length, length)

	cnt := 0
	for i := 0; i < length; i++ {
		if cnt == 0 {
			str[cnt] = input[i]
			cnt++
			continue
		}
		if str[cnt-1] == input[i] {
			cnt--
		} else {
			str[cnt] = input[i]
			cnt++
		}

	}
	return string(str[:cnt])
}
