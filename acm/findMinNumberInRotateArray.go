package acm

/**
顺序数组颠倒部分
查找最小数

 */
func minNumberInRotateArray( rotateArray []int ) int {
	// write code here
	if len(rotateArray) == 0 {
		return 0
	}
	var start int = 0
	var end int = len(rotateArray) -1
	var mid int = end/1
	for start < end {
		mid = (start + end) / 2
		if rotateArray[mid] <= rotateArray[end] {
			end = mid
		}  else {
			start = mid + 1
		}
	}
	return rotateArray[start]
}
