package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if(x < 0) {
		return false
	}

	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i < l/2; i++{
		if(s[i] == s[l-i-1]) {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {

	var digit int

	for {
		fmt.Scanln(&digit)
		fmt.Println(isPalindrome(digit))
	}
}
