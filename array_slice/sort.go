package arrayslice

import (
	"fmt"
	"sort"
)

func SortAsc(input []int) []int {
	n := len(input)
	for i := 0; i <= n-1; i++ {
		for j := i + 1; j < n; j++ {
			if input[i] > input[j] {
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
			}
		}
	}
	return input
}

func SortDsc(input []int) []int {
	n := len(input)
	for i := 0; i <= n-1; i++ {
		for j := i + 1; j < n; j++ {
			if input[i] < input[j] {
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
			}
		}
	}
	return input
}

func SortWithInbuiltSort(input []int) []int {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	return input
}

func SortAscWithSort(input []int) {
	sort.Ints(input)
	fmt.Println("sort ints", input)
}

func SortStringAscWithSort() {
	str := []string{"s", "d", "c", "b", "a"}
	sort.Strings(str)
	fmt.Println("sort Strings", str)
}

func SortFirstHalfAndSecondHalf1(input []int) {
	sort.Ints(input)
	var output []int
	n := len(input)
	for i := 0; i < n/2; i++ {
		output = append(output, input[i])
	}
	for j := n - 1; j >= n/2; j-- {
		output = append(output, input[j])
	}
	fmt.Println("SortFirstHalfAndSecondHalf1", output)
}

func SortFirstHalfAndSecondHalf2(input []int) {
	n := len(input)
	for i := 0; i < n-1; i++ {
		//Asc
		for j := 0; j < n/2; j++ {
			if input[j] > input[j+1] {
				temp := input[j]
				input[j] = input[j+1]
				input[j+1] = temp
			}
		}

		//Dsc
		for j := n / 2; j < n-1; j++ {
			if input[j] < input[j+1] {
				temp := input[j]
				input[j] = input[j+1]
				input[j+1] = temp
			}
		}
	}

	fmt.Println("SortFirstHalfAndSecondHalf2", input)
}
