package stringprograms

import "fmt"

func LenghtOfString1(run []rune) int {
	var count int = 0
	for _, rune := range run {
		fmt.Printf("%c", rune)
		count++
	}
	return count
}

func LenghtOfString2(run []rune) int {
	var count int = 0
	for i := 0; i < len(run); i++ {
		count++
	}
	return count
}
