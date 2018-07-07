package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	max_end := nums[0]
	max_len := nums[0]
	for i := 1; i < len(nums); i++ {
		max_end = int(math.Max(float64(nums[i]), float64(max_end+nums[i])))
		max_len = int(math.Max(float64(max_len), float64(max_end)))
	}
	return max_len
}

func main() {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(a))
}