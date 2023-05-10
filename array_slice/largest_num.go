package arrayslice

func LargestNum(arr []int) int {
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func SmallestNum(arr []int) int {
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func LargestAndSmallestNum(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}
