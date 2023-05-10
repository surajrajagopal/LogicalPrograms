package stringprograms

func PalindromeOrNot1(str string) bool {
	reverseString := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		reverseString = append(reverseString, str[i])
	}
	return str == string(reverseString)
}

func PalindromeOrNot2(str string) bool {
	var result string
	for _, v := range str {
		result = string(v) + result
	}
	return str == result
}

func PalindromeOrNot3(str string) bool {
	for i := 0; i < (len(str)/2)+1; i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
