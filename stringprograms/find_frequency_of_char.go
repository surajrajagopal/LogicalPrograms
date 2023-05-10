package stringprograms

func FrequencyOfChar(str string) map[string]int {
	frequencyCount := make(map[string]int)
	for _, v := range str {
		if frequencyCount[string(v)] == 0 {
			frequencyCount[string(v)] = 1
		} else {
			frequencyCount[string(v)] = frequencyCount[string(v)] + 1
		}
	}
	return frequencyCount
}
