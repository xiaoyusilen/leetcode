package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	index3 := m + n - 1
	for {
		if index1 < 0 || index2 < 0 {
			break
		}
		if nums1[index1] > nums2[index2] {
			nums1[index3] = nums1[index1]
			index1--
			index3--
		} else {
			nums1[index3] = nums2[index2]
			index2--
			index3--
		}
	}
	for ; index2 >= 0; {
		nums1[index3] = nums2[index2]
		index2--
		index3--
	}
	fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{0}, 0, []int{1}, 1)
}