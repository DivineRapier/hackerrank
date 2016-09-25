package tutorials

import (
	"fmt"
	"regexp"
	"sort"
)

func regex_patterns_databases() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var T int
	var name, email string
	var ret []string

	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%s%s", &name, &email)
		if match, _ := regexp.Match(".*@gmail\\.com", []byte(email)); match {
			ret = append(ret, name)
		}
	}
	sort.Strings(ret)
	for _, v := range ret {
		fmt.Println(v)
	}
}
