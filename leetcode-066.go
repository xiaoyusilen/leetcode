package main

import "fmt"

func plusOne(digits []int) []int {
	digits[len(digits) - 1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if(digits[i] <= 9) {
			return digits
		} else {
			digits[i] = 0
			if(i - 1 >= 0) {
				digits[i-1]++
			} else {
				digits = append([]int{1}, digits[:]...)
			}
		}
	}
	return digits
}

func main() {
	a := []int{8, 9, 9, 9}
	fmt.Println(plusOne(a))
	a = []int{1, 2, 3}
	fmt.Println(plusOne(a))
	a = []int{9, 9, 9, 9}
	fmt.Println(plusOne(a))
}