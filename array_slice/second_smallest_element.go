package arrayslice

import (
	"fmt"
	"math"
)

func Second_Smallest_Element1(input []int) int {
	min := math.MaxInt64
	second_min := math.MaxInt64

	n := len(input)
	for i := 0; i < n; i++ {
		if input[i] < min {
			min = input[i]
		}
	}

	fmt.Println("min", min)
	for i := 0; i < n; i++ {
		if input[i] != min && input[i] < second_min {
			second_min = input[i]
		}
	}
	return second_min
}

func Second_Smallest_Element2(input []int) int {
	min := math.MaxInt64
	second_min := math.MaxInt64

	n := len(input)
	for i := 0; i < n; i++ {
		if input[i] < min {
			second_min = min
			min = input[i]
		}
	}
	for i := 0; i < n; i++ {
		if input[i] != min && input[i] < second_min {
			second_min = input[i]
		}
	}
	return second_min
}
