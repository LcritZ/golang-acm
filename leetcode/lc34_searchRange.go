package leetcode

/**
二分查找
找到最左和最右的边界

 */
func SearchRange(nums []int, target int) []int {
    res := []int{-1, -1}
    if len(nums) == 0 {
        return res
    }

    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            right = mid-1
        } else {
            left = mid+1
        }
    }

    if left >= len(nums) || nums[left] != target  {
        return res
    } else {
        res[0] = left
        res[1] = left
    }

    if left < len(nums)-1 && nums[left+1] != target {
        return res
    }

    left, right = 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            left = mid+1
        } else {
            right = mid-1
        }
    }

    if right > res[1] {
        res[1] = right
    }
    return res
}
