package basicprograms

func SumOfNNaturalNumber(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func FromulaSumOfNNaturalNumber(num int) int {
	sum := num * (num + 1) / 2
	return sum
}

func RecursionSumOfNNaturalNumber(num int) int {
	if num == 0 {
		return num
	} else {
		return num + RecursionSumOfNNaturalNumber(num-1)
	}
}
