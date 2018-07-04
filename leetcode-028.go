package main

import (
	"strings"
	"fmt"
)

func strStr(haystack string, needle string) int {
	// warning: judge len(needle), not haystack
	if len(needle) == 0 {
		return 0
	}
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}