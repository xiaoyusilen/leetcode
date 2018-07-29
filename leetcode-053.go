package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 题解：http://xiaoyu.world/Algo/leetcode-014&053/

func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = getMax(nums[i], sum + nums[i])
		max = getMax(sum, max)
	}
	return max
}

func main() {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(a))
}