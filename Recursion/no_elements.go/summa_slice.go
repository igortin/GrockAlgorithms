package main

import (
	"fmt"
)

func main() {
	n := []int{1, 2, 7}
	r := summa(n)

	fmt.Println(r)
}

func summa(n []int) int {
	if len(n) == 0 {
		return 0
	}
	value := n[0]
	n = n[1:]
	return value + summa(n)
}
