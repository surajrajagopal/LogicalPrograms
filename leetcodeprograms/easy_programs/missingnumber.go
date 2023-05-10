package leetcodeprograms

import arrayslice "logic/array_slice"

func MissingNumber(arr []int) int {
	n := len(arr)
	acutalSum := n * (n + 1) / 2
	numsSum := arrayslice.SumofElements(arr)
	return acutalSum - numsSum
}

// //var dummy *Value = nil
// var _ Calculate = (*Value)(nil)
