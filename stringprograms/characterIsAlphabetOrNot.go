package stringprograms

import "unicode"

func AlphabetOrNot1(char []rune) (alpha []rune) {
	for i := 0; i < len(char); i++ {
		if unicode.IsLetter(char[i]) {
			alpha = append(alpha, char[i])
		} else {
			continue
		}
	}
	return
}

func AlphabetOrNot2(char []byte) (alpha []byte) {
	for i := 0; i < len(char); i++ {
		if char[i] >= 65 && char[i] <= 90 || char[i] >= 97 && char[i] <= 122 {
			alpha = append(alpha, char[i])
		} else {
			continue
		}
	}
	return
}
