package main

import (
	"fmt"
	"math/rand"
)

var (
	slicelength = 10
)

func createList() []int {
	slice := make([]int, 0)
	for i := 0; i < slicelength; i++ {
		slice = append(slice, rand.Intn(10))
	}
	return slice
}

func findMin(s *[]int) int {
	pos := 0

	minValue := (*s)[pos]
	for i := 0; i < len(*s); i++ {
		if (*s)[i] < minValue {
			minValue = (*s)[i]
			pos = i
		}
	}
	*s = append((*s)[:pos], (*s)[pos+1:]...)
	return minValue
}

func main() {
	s := createList()

	fmt.Println("random slice:", s)
	resSlice := make([]int, 0)
	for i := 0; i < slicelength; i++ {
		v := findMin(&s)
		resSlice = append(resSlice, v)
	}

	fmt.Println("result slice:", resSlice)
}
