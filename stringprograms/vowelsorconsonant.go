package stringprograms

func isLowerCaseVowels(run rune) bool {
	return run == 'a' || run == 'e' || run == 'i' || run == 'o' || run == 'u'
}

func isUpperCaseVowels(run rune) bool {
	return run == 'A' || run == 'E' || run == 'I' || run == 'O' || run == 'U'
}
func VowelsOrConsonants(run []rune) (vowels []rune, Consonants []rune) {
	for _, rune := range run {
		if isLowerCaseVowels(rune) || isUpperCaseVowels(rune) {
			vowels = append(vowels, rune)
		} else {
			Consonants = append(Consonants, rune)
		}
	}
	return
}
