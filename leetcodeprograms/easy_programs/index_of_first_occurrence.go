package leetcodeprograms

import (
	"strings"
)

func StrStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}

	for i := 0; i < (len(haystack)-len(needle))+1; i++ {
		builder := strings.Builder{}
		for j := i; j < i+len(needle); j++ {
			builder.WriteRune(rune(haystack[j]))
		}
		if builder.String() == needle {
			return i
		}
	}

	return -1
}

func Index_Of_First_Occurence(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
