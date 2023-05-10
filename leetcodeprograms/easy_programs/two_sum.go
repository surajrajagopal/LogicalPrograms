package leetcodeprograms

func TwoSum(arr []int, target int) []int {
	var pairs []int
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		val := target - arr[i]
		if _, ok := m[val]; ok {
			p := []int{m[val], i}
			pairs = append(pairs, p...)
		}
		m[arr[i]] = i
	}
	return pairs
}
