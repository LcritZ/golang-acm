package basic

import (
	"fmt"
)

func BubbleSort(input []int) []int {
	if len(input) == 0 {
		return input
	}
	for i := 0; i < len(input); i++ {
		flag := false
		for j := 0; j < len(input)-1; j++ {
			if input[j+1] < input[j] {
				flag = true
				input[j+1], input[j] = input[j], input[j+1]
				fmt.Println("--", input)
			}
		}
		fmt.Println("++", i)
		if !flag {
			fmt.Println(i)
			break
		}
	}
	return input
}

func SelectSort(input []int) []int {
	if len(input) == 0 {
		return input
	}
	for i := 0; i < len(input); i++ {
		index := i
		temp := input[i]
		for j := i + 1; j < len(input); j++ {
			if input[j] < temp {
				temp = input[j]
				index = j
			}
		}
		input[i], input[index] = input[index], input[i]
	}

	return input
}

func InsertSort(input []int) []int {
	if len(input) == 0 {
		return input
	}

	for i := 1; i< len(input); i++ {
		temp := input[i]
		j := i-1
		if temp < input[j] {
			for ; j >=0 && temp < input[j]; j--  {
				input[j+1] = input[j]
			}
			input[j+1] = temp
		}

	}

	return input
}

func QuickSort(input []int, low, high int) {
	if low < high {
		mid := partion(input, low, high)
		fmt.Println(input[mid])
		QuickSort(input, low, mid-1)
		QuickSort(input, mid+1, high)
	}
}

func partion(arr []int, low, high int) int {
	temp := arr[low]
	for low < high {
		for low<high && arr[high] > temp {
			high--
		}
		arr[low] = arr[high]
		for low<high && arr[low] <= temp {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = temp
	return low
}

func QuickSort3(nums []int, low, high int)  {
	for low < high {
		mid := helper3(nums, low, high)
		QuickSort(nums, low, mid-1)
		QuickSort(nums, mid+1, high)
	}
}

func helper3(nums []int, low, high int) int {
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
	return low
}

func MergeSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := []int{}
	var sortHelper func(nums []int, low, high int) []int
	sortHelper = func(nums []int, low, high int) []int {
		if len(nums) == 0 {
			return []int{}
		}
		if low == high {
			return nums[low:low+1]
		}

		mid := (low+high)>>1
		n1 := sortHelper(nums, low, mid)
		n2 := sortHelper(nums, mid+1, high)
		temp := []int{}
		i := 0
		j := 0
		for i < len(n1) && j < len(n2) {
			if n1[i] <= n2[j] {
				temp = append(temp, n1[i])
				i++
			} else {
				temp = append(temp, n2[j])
				j++
			}
		}

		if i < len(n1) {
			temp = append(temp, n1[i:]...)
		} else {
			temp = append(temp, n2[j:]...)
		}
		if len(n1)+len(n2) == len(nums) {
			res = make([]int, len(nums))
			copy(res, temp)
		}

		return temp
	}

	sortHelper(nums, 0, len(nums)-1)
	return res
}
