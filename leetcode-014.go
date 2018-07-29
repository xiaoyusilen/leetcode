package main

import (
	"fmt"
	"strings"
)

// 题解：http://xiaoyu.world/Algo/leetcode-014&053/
func longestCommonPrefix(strs []string) string {
	if(len(strs) == 0) {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for ; strings.Index(strs[i], prefix) != 0; {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	a := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(a))
	a = []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(a))
}