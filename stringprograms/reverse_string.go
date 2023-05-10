package stringprograms

func ReverseString1(run []rune) []rune {
	reverseString := []rune{}
	for i := len(run) - 1; i >= 0; i-- {
		reverseString = append(reverseString, run[i])
	}
	return reverseString
}

func ReverseString2(run string) (result string) {
	for _, v := range run {
		result = string(v) + result
	}
	return
}
