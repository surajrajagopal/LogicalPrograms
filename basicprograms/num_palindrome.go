package basicprograms

func PalindromNum(n int) bool {
	temp := n
	reverse := 0
	for n != 0 {
		rem := n % 10
		reverse = reverse*10 + rem
		n /= 10
	}
	return temp == reverse
}
