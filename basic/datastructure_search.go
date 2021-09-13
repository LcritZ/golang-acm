package basic

import "fmt"

func BinarySearchRecursion(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}
	//if len(arr) ==1 {
	//	return arr[0] == target
	//}
	mid := len(arr)/2
	if target == arr[mid] {
		return true
	}
	if target > arr[mid] {
		return BinarySearchRecursion(arr[mid+1:], target)
	} else {
		return BinarySearchRecursion(arr[:mid], target)
	}
}

func BinarySearchRecursion2(arr []int, low, high, target int) bool {
	if low >= high {
		//fmt.Println(arr[low], arr[high])
		return false
	}
	if len(arr) == 0 {
		return false
	}
	//if len(arr) ==1 {
	//	return arr[0] == target
	//}
	mid := low + (high-low)/2
	if target == arr[mid] {
		return true
	}
	fmt.Println(mid, low, high)
	if target > arr[mid] {
		return BinarySearchRecursion2(arr, mid+1, high, target)
	} else {
		return BinarySearchRecursion2(arr, low, mid-1, target)
	}
}

func BinarySearchRecursion3(arr []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if target == arr[mid] {
		return mid
	}
	if target > arr[mid] {
		return BinarySearchRecursion3(arr, mid+1, high, target)
	} else {
		return BinarySearchRecursion3(arr, low, mid-1, target)
	}
}

func BinarySearchNoRecursion(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}
	low := 0
	high := len(arr)-1
	for low < high {
		mid := low + (high-low)/2
		fmt.Println(arr[low:high+1])
		if target == arr[mid] {
			return true
		}
		if target > arr[mid] {
			low = mid
		} else {
			high = mid
		}
	}
	return false
}

func BinarySearchNoRecursion2(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low := 0
	high := len(arr)-1
	for low < high {
		mid := low + (high-low)/2
		fmt.Println(arr[low:high+1])
		if target == arr[mid] {
			return mid
		}
		if target > arr[mid] {
			low = mid
		} else {
			high = mid
		}
	}
	return -1
}