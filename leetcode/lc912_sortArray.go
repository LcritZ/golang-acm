package leetcode

func SortArray(nums []int) []int {
    QuickSort(nums, 0, len(nums)-1)
    return nums
}

func QuickSort(nums []int, low, high int) {
    if low < high {
        mid := Helper912(nums, low, high)
        QuickSort(nums, low, mid-1)
        QuickSort(nums, mid+1, high)
    }
}

func Helper912(nums []int, low, high int) int {
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
