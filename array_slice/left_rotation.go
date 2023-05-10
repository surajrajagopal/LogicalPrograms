package arrayslice

// Input arr:=[3,1,0,8] , RotateBy = 2, output := [0,8,3,1]

// Rotations take place : [1,0,8,3] , [0,8,3,1]
func LeftRotate(arr []int, count int) []int {
	for i := 0; i < count; i++ {
		var j int
		a := arr[0]
		for j = 1; j < len(arr); j++ {
			arr[j-1] = arr[j]
		}
		arr[j-1] = a
	}
	return arr
}
