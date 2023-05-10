package stringprograms

func RemoveBracket(str string) string {
	removedBrackets := ""
	for _, v := range str {
		if string(v) != "(" && string(v) != ")" {
			removedBrackets += string(v)
		}
	}
	return removedBrackets
}
