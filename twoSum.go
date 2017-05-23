// author by @xiaoyusilen

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				break
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Print(result)
}
