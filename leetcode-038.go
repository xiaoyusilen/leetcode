package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if(n==1) {
		return "1"
	}
	s := countAndSay(n-1) + "*"
	var str string
	count := 1
	for i := 0; i < len(s) - 1; i++ {
		if(s[i] == s[i+1]) {
			count++
		} else {
			str = str + strconv.Itoa(count) + string(s[i])
			count = 1
		}
	}
	return str
}

func main() {
	fmt.Println(countAndSay(4))
}
