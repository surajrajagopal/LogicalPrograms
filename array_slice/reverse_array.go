package arrayslice

import (
	"fmt"
	"reflect"
	"sort"
)

func Reverse1(input []int) []int {
	var output []int

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}

func Reverse2(input []int) []int {
	inputLen := len(input)
	output := make([]int, len(input))

	for i, n := range input {
		j := inputLen - i - 1
		output[j] = n
	}
	return output
}

func Reverse3(input []int) []int {
	inputLen := len(input)
	inputMid := len(input) / 2

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1
		input[i], input[j] = input[j], input[i]

	}
	return input
}

func Reverse4(input interface{}) []int32 {
	inputLen := reflect.ValueOf(input).Len()
	inputMid := inputLen / 2
	inputSwap := reflect.Swapper(input)

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		inputSwap(i, j)
	}
	return input.([]int32)
}

func Reverse5(input []int) []int {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func Reverse6(input []int) []int {
	start := 0
	end := len(input) - 1

	for end < start {
		input[start], input[end] = input[end], input[start]
	}
	return input
}

func Reverse7(input []int) []int {
	start := 0
	end := len(input) - 1

	for start < end {
		temp := input[start]
		input[start] = input[end]
		input[end] = temp
		start++
		end--
	}
	return input
}

func Reverse8(input []int) {
	sort.Sort((sort.Reverse(sort.IntSlice(input))))
	fmt.Println("inbuilt sort", input)
}
