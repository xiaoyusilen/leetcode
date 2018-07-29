package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen % 2 == 0 {
		return findKthNumberInTwoArray(nums1, nums2, 0, 0, totalLen/2)/2 +
			findKthNumberInTwoArray(nums1, nums2, 0, 0, totalLen/2 + 1)/2
	} else {
		return findKthNumberInTwoArray(nums1, nums2, 0, 0, totalLen/2 + 1)
	}
}

func findKthNumberInTwoArray(nums1, nums2 []int, p, q, k int) float64 {
	if p >= len(nums1) {
		return float64(nums2[q + k - 1])
	}
	if q >= len(nums2) {
		return float64(nums1[p + k - 1])
	}
	if k == 1 {
		return math.Min(float64(nums1[p]), float64(nums2[q]))
	}
	m := k/2
	n := k - m
	a := math.MaxInt32
	b := a
	if p + m - 1 < len(nums1) {
		a = nums1[p + m - 1]
	}
	if q + n - 1 < len(nums2) {
		b = nums2[q + n - 1]
	}
	if a < b {
		return findKthNumberInTwoArray(nums1, nums2, p + m, q, k - m)
	} else {
		return findKthNumberInTwoArray(nums1, nums2, p, q + n, k - n)
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3, 4}))
}
