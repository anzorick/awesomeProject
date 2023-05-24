package main

import (
	"fmt"
	"math"
)

const size = 10

func twoSmallesnumbers(input [size]int) (int, int) {
	min := input[0]

	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}

	secondMin := math.MaxInt
	for i := 0; i < len(input); i++ {
		if input[i] < secondMin && input[i] > min {
			secondMin = input[i]
		}
	}
	return min, secondMin
}

func NewtwoSmallesNumbers(input [size]int) (int, int) {
	min := math.MaxInt
	secondMin := math.MaxInt

	for i := 0; i < len(input); i++ {
		if input[i] < min {
			secondMin = min
			min = input[i]
		} else if input[i] < secondMin {
			secondMin = input[i]
		}
	}
	return min, secondMin
}
func main() {
	numbers := [size]int{10, 20, 30, 40, 50, 90, 80, 70, 2, 60}
	fmt.Println(twoSmallesnumbers(numbers))
	fmt.Println(NewtwoSmallesNumbers(numbers))
}
