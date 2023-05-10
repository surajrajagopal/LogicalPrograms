package stringprograms

import "unicode"

func IsAlphabet(r rune) bool {
	return (unicode.ToUpper(r) >= 'A' && unicode.ToUpper(r) <= 'Z') ||
		(unicode.ToLower(r) >= 'a' && unicode.ToLower(r) <= 'z')
}
func RemoveSpecialAndNumCharacters(str string) string {
	var runs []rune
	for _, v := range str {
		if IsAlphabet(v) {
			runs = append(runs, v)
		} else {
			continue
		}
	}
	return string(runs)
}
