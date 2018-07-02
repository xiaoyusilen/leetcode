package main

import (
	"strings"
	"fmt"
)

func isValid(s string) bool {
	for {
		l := len(s)
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		if(len(s) == l) {
			break
		}
	}

	return len(s) == 0
}

func main() {
	var s string

	for {
		fmt.Scanln(&s)
		fmt.Println(isValid(s))
	}
}
