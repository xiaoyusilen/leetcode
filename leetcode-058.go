package main

import (
	"strings"
	"fmt"
)

func lengthOfLastWord(s string) int {
	str := strings.Split(Reverse(s), " ")
	for _, key := range str {
		if(!strings.EqualFold(key, "")) {
			return len(key)
		}
	}
	return 0
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord(" a "))
	fmt.Println(lengthOfLastWord("        "))
}
