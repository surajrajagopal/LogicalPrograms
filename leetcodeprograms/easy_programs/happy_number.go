package leetcodeprograms

func IsHappy(num int) bool {
	for num != 1 {
		records := make(map[int]int)
		records[num] = num
		num = getSqOfDigits(num)
		for _, previous := range records {
			if num == previous {
				return false
			}
		}
	}
	return true
}

func getSqOfDigits(num int) int {
	sqOfDigits := 0
	for num != 0 {
		rem := num % 10
		sqOfDigits += rem * rem
		num /= 10
	}
	return sqOfDigits
}
