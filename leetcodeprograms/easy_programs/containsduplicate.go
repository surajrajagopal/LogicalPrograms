package leetcodeprograms

func ContainsDuplicate(nums []int) bool {
	consdup := make(map[int]int)
	for _, v := range nums {
		consdup[v]++
	}
	for _, v := range consdup {
		if v > 1 {
			return true
		}
	}
	return false
}

func ContainsDuplicate1(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
