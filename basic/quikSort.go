package basic

func QuickSort2(arr []int, low, high int) {
	if low < high {
		midIndex := helper(arr, low, high)
		QuickSort(arr, low, midIndex-1)
		QuickSort(arr, midIndex+1, high)
	}
}

func helper(arr []int, low, high int) int {
	mid := arr[low]
	temp := arr[low]
	for low < high {
		for low < high && arr[high] > mid {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= mid {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = temp
	return low
}
