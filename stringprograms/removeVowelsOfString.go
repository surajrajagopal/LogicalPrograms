package stringprograms

import "regexp"

func RemoveVowelsOfString(str string) string {
	reg := regexp.MustCompile(`[aeiouAEIOU]`)
	repVowels := reg.ReplaceAllString(str, "")
	return repVowels
}
