package basic

/**
是否是出栈顺序
 */
func IsPopOrder( pushV []int ,  popV []int ) bool {
	// write code here
	length := len(pushV)
	if length == 0 {
		return false
	}

	var arr []int
	var j int = 0
	for i := 0; i < length; i++ {
		arr = append(arr, pushV[i])
		for len(arr) != 0 && arr[len(arr)-1] == popV[j]  {
			arr = arr[:len(arr)-1]
			j ++
		}
	}

	if len(arr) != 0 {
		return false
	}
	return true
}
