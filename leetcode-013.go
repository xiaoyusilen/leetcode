package main

import "fmt"

func romanToInt(s string) int {
	var r = []rune(s)
	l := len(r)
	if(l == 0) {
		return 0
	}

	m := map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}

	d := m[string(r[l-1])]

	for i := l-2; i >= 0; i-- {
		if(m[string(r[i])] >= m[string(r[i+1])]) {
			d = d + m[string(r[i])]
		} else {
			d = d - m[string(r[i])]
		}
	}
	return d
}

func main() {
	var str string

	for {
		fmt.Scanln(&str)
		fmt.Println(romanToInt(str))
	}
}
