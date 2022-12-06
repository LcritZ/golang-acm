package leetcode

import "fmt"

func Search33(arr []int, target int) int {
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
		if arr[0] <= arr[mid] {
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


/**
4,5,6,7,0,1,2

与第0个对比确定那边是正序

target在正序中正确调整，否则其他各种情况都调整另外一边

注意边界条件的判断

 */
func Search33_2(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	//根据范围，判断哪一边是顺序的，然后在这一段，判断target是否在这个区间，不然的话，就找另外一边 一直都是和0对比，判断是否在顺序
	for left <= right {
		mid := left + (right-left)/2
		fmt.Println(nums[left], nums[right], nums[mid])
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			if target >= nums[0] && target < nums[mid] {
				right = mid-1
			} else {
				left = mid+1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid+1
			} else {
				right = mid-1
			}
		}
	}

	return -1
}

func Search33_3(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			if target >= nums[0] && target < nums[mid] {
				right = mid-1
			} else {
				left = mid+1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid+1
			} else {
				right = mid-1
			}
		}
	}

	return -1

}
