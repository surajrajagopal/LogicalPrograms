package stringprograms

func RemoveDuplicates(str string) map[string]int {
	frequencyCount := make(map[string]int)
	for _, v := range str {
		if _, ok := frequencyCount[string(v)]; !ok {
			frequencyCount[string(v)] = 1
		} else {
			frequencyCount[string(v)]++
		}
	}

	CopiedMap := make(map[string]int)
	for k, v := range frequencyCount {
		if v == 1 {
			CopiedMap[k] = v
		}
	}
	return CopiedMap
}
