package main

import (
	"fmt"
)

func main() {
	n := []int{5, 1, 7}
	fmt.Println(qsortFunc(n))
}

func qsortFunc(n []int) []int {
	if len(n) < 2 {
		return n
	} else {
		pivot := n[0]
		n = n[1:]
		lessPart := make([]int, 0)
		biggerPart := make([]int, 0)
		for i := 0; i < len(n); i++ {
			if n[i] < pivot {
				lessPart = append(lessPart, n[i])
			}
			if n[i] >= pivot {
				biggerPart = append(biggerPart, n[i])
			}
		}
		lessArr := qsortFunc(lessPart)
		biggerArr := qsortFunc(biggerPart)
		lessArr = append(lessArr, pivot)
		lessArr = append(lessArr, biggerArr...)
		return lessArr
	}
}
