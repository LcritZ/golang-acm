package basic

/**
二维数组查找
 */
func FindNum(arr [][]int, target int) bool {
	if len(arr) == 0 {
		return false
	}

	i := 0
	j := len(arr)-1
	for i < len(arr[0]) -1 && j > 0 {
		if target > arr[j][i] {
			i++
		} else if target == arr[j][i]{
			return true
		} else {
			j--
		}
	}

	return false
}
