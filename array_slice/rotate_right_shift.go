package arrayslice

import "fmt"

func RightRotatationCopy() {
	num := 4 % 8
	old_array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	new_array := make([]int, len(old_array))

	copy(new_array[:num], old_array[len(old_array)-num:])
	copy(new_array[num:], old_array[:len(old_array)-num])
	copy(old_array, new_array)
	fmt.Println(old_array)
}
