package basic

import "fmt"

func QuickSort2(arr []int, low, high int) {
	if low < high {
		midIndex := helper(arr, low, high)
		QuickSort2(arr, low, midIndex-1)
		QuickSort2(arr, midIndex+1, high)
	}
}

func helper(nums []int, low, high int) int {
	//mid := arr[low]
	////temp := arr[low]
	//for low < high {
	//	for low < high && arr[high] > mid {
	//		high--
	//	}
	//	arr[low] = arr[high]
	//	for low < high && arr[low] <= mid {
	//		low++
	//	}
	//	arr[high] = arr[low]
	//}
	//arr[low] = mid
	//return low
	mid := nums[low]
	for low < high {
		for low < high && nums[high] > mid {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= mid {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = mid
	fmt.Println(nums, low)
	return low
}
