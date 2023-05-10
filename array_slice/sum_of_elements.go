package arrayslice

func SumofElements(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
