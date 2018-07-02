package main

import "fmt"
import (
	"math"
)

const MAX = 2147483647

func reverse(x int) int {
	s := 0
	for {
		s = s * 10
		s = s + x%10
		x = x/10
		if(x == 0) {
			break
		}
	}
	if(math.Abs(float64(s)) > MAX) {
		return 0
	}
	return s
}

func main() {

	var digit int

	for {
		fmt.Scanln(&digit)
		fmt.Println(reverse(digit))
	}
}
