package leetcode

import "fmt"

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums) - k + 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize/2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i * 2 + 1, i * 2 + 2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}






func getK(arr []int, k int) (int, bool) {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return 0, false
	}
	sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	index := len(arr) - k
	return arr[index], true
}

func sort(arr []int, low, high int) {
	if low < high {
		mid := helper215(arr, low, high)
		sort(arr, low, mid-1)
		sort(arr, mid+1, high)
	}
}

func helper215(arr []int, low, high int) int {
	temp := arr[low]
	mid := arr[low]
	for low < high {
		for low < high && arr[high] > mid {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= mid {
			low ++
		}
		arr[high] = arr[low]
	}
	arr[low] = temp
	return low
}

func findKthLargest2(nums []int, k int) int {
	if k <= 0 || len(nums) == 0 {
		return 0
	}

	// 分区操作，返回轴点索引下标
	// 在数组nums的子区间 [left, right] 执行 partition 操作，返回 nums[left]（选取的第一个作为pivot） 排序以后应该在的位置
	partition := func(nums []int, left, right int) int {
		pivot := nums[left]
		j := left
		for i := left + 1; i <= right; i++ {
			// 小于 pivot 的元素都被交换到前面
			if nums[i] < pivot {
				j++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		// 在之前遍历的过程中，满足 [left+1, j] < pivot，并且 (j, i] >= pivot
		nums[j], nums[left] = nums[left], nums[j]
		// 交换以后 [left, j-1] < pivot, nums[j] = pivot, [j+1, right] >= pivot
		return j
	}

	n := len(nums)
	// 转换一下，第 k 大元素的索引是 n - k
	left, right, target := 0, n-1, n-k
	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

func getK2(arr []int, k int) int {
	if k <= 0 || len(arr) == 0 {
		return 0
	}
	length := len(arr)
	left, right, target := 0, length-1, length-k
	for {
		index := helper215(arr, left, right)
		if index == target {
			return arr[index]
		} else if index < target {
			left = index+1
		} else {
			right = index-1
		}
	}



}