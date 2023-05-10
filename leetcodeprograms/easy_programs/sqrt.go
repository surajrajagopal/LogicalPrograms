package leetcodeprograms

import "fmt"

func MySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) / 2
		//	fmt.Println(mid)
		if mid*mid > x {
			r = mid - 1
			fmt.Println(r)
		} else {
			l = mid
		}
	}
	return l
}
