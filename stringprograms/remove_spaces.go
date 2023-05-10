package stringprograms

import "strings"

func RemoveSpaces1(str string) string {
	s := strings.ReplaceAll(str, " ", "")
	return s
}

func RemoveSpaces2(str string) string {
	removedSpace := ""
	for _, v := range str {
		if string(v) != " " {
			removedSpace += string(v)
		}
	}
	return removedSpace
}
