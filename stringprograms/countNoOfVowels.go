package stringprograms

import "unicode"

func isVowels(r rune) bool {
	return unicode.ToUpper(r) == 'A' ||
		unicode.ToUpper(r) == 'E' ||
		unicode.ToUpper(r) == 'I' ||
		unicode.ToUpper(r) == 'O' ||
		unicode.ToUpper(r) == 'U'
}
func CountNoOfVowels(run []rune) int {
	count := 0
	for _, rune := range run {
		if isVowels(rune) {
			count++
		}
	}
	return count
}
