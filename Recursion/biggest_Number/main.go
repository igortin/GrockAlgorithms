package main

import (
	"fmt"
)

func main() {
	n := []int{1, 7, 7, 7, 2, 5}
	res := biggestNumber(n)

	fmt.Println(res)

}

func biggestNumber(n []int) int {

	if len(n) == 0 {
		return 0
	}
	res := n[0]
	n = n[1:]

	tmp := biggestNumber(n)
	if res < tmp {
		res = tmp
	}

	return res
}
