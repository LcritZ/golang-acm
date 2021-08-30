package leetcode

func Search(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	low := 0
	high := len(arr)-1
	for low < high {
		mid := low + (high-low)/2
		if target == arr[mid] {
			return mid
		}
		if arr[0] <= mid {
			if target >= arr[0] && target <= arr[mid] {
				high = mid-1
			} else {
				low = mid+1
			}
		} else {
			if target >= arr[mid] && target <= arr[len(arr)-1] {
				low = mid+1
			} else {
				high = mid-1
			}
		}
	}

	return -1
}
