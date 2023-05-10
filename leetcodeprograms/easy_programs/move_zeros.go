package leetcodeprograms

//Input: [0,1,0,3,12]
//Output: [1,3,12,0,0]

func MoveZeroes(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			j++
		}
	}
	return arr
}
