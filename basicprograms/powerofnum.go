package basicprograms

func PowerOfNum(n int) int {
	expo := 2
	power := 1
	for i := 0; i <= expo; i++ {
		power = power * n
	}
	return power
}
