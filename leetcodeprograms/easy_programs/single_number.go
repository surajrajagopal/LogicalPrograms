package leetcodeprograms

func SingleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
		if m[v] > 1 {
			delete(m, v)
		}
	}
	for k := range m {
		return k
	}
	return 0
}
