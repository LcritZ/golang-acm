package acm

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
