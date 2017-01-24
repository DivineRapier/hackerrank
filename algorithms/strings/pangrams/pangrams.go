package strings

import "strings"

func pangrams(s string) string {
	m := make(map[byte]bool)
	l := len(s)
	str := strings.ToLower(s)

	for i := 0; i < l; i++ {
		if str[i] == ' ' {
			continue
		}
		m[str[i]] = true
	}
	if len(m) == 26 {
		return "pangram"
	}
	return "not pangram"
}
