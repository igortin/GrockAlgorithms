package main

import (
	"fmt"
	"strconv"
)

var (
	sliceLegth = 100
)

func fillArray(array *[]int) {
	for i := 0; i < sliceLegth; i++ {
		*array = append(*array, i)
	}
}

func main() {
	var item string
	inputList := make([]int, 0)
	fillArray(&inputList)
	fmt.Println(inputList)
	fmt.Print("Input item: ")
	fmt.Scanln(&item)
	val, err := strconv.Atoi(item)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Answer index:", binarySearch(&inputList, val))
}

func binarySearch(array *[]int, item int) int {
	i := 0
	j := len(*array) - 1
	var middlePos int

	for i <= j {
		if item > (*array)[sliceLegth-1] {
			break
		}
		if len(*array)%2 == 1 {
			middlePos = (j + i + 1) / 2
		} else {
			middlePos = (j + i) / 2
		}
		fmt.Println(i, middlePos, j)
		if (*array)[middlePos] > item {
			j = middlePos - 1 // исключаем *array[middlePos]
		}
		if (*array)[middlePos] < item {
			i = middlePos + 1 // исключаем *array[middlePos]
		}
		if (*array)[middlePos] == item {
			return middlePos
		}
	}
	return -1
}
