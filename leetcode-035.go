package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if(nums[i] < target) {
			if(i == len(nums)-1 || nums[i+1] > target) {
				return i + 1
			}
		} else if(nums[i] == target) {
			return i
		}
	}
	return 0
}

func main() {
	arr := []int{1,3,5,6}
	fmt.Println(searchInsert(arr, 0))
}
