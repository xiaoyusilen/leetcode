package main

import "fmt"

var m = map[int]int{
	0: 1,
	1: 1,
}

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		res := m[n]
		// 优化点，记录每一次结果，避免多次重复递归调用
		if res == 0 {
			res = climbStairs(n - 2) + climbStairs(n - 1)
			m[n] = res
		}

		return res
	}
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(44))

}
