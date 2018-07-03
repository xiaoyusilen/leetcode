package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if(len(nums) == 0) {
		return 0
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if(nums[i] != nums[count]) {
			count++
			nums[count] = nums[i];
		}
	}
	return count+1
}

func main() {

	arr := []int{1,1,2}
	for i := 0; i < removeDuplicates(arr); i++ {
		fmt.Println(arr[i])
	}
}

