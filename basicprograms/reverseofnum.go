package basicprograms

func ReverseOfDigits(n int) int {
	reverse := 0
	for n != 0 {
		rem := n % 10
		reverse = reverse*10 + rem
		n /= 10
	}
	return reverse
}
