package basicprograms

func ArmstrongNum(n int) bool {
	temp := n
	sum := 0
	for n != 0 {
		rem := n % 10
		sum += rem * rem * rem
		n /= 10
	}
	return temp == sum
}
