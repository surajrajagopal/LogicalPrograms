package arrayslice

// Input arr:=[3,1,0,8] , RotateBy = 2, output := [8,0,1,3]
func RightRotate(arr []int, rotateBy int) []int {
	for i := 0; i < rotateBy; i++ {
		var j int
		last := arr[len(arr)-1]
		for j = len(arr) - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = last
	}
	return arr
}
