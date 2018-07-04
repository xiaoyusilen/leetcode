package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	if(len(nums) == 0) {
		return 0
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if(nums[i] != val) {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func main() {
	arr := []int{0,1,2,2,3,0,4,2}
	l := removeElement(arr, 2)
	for i := 0; i < l; i++ {
		fmt.Println(arr[i])
	}
}

