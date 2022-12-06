package leetcode

func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }



    left, right := 0, len(nums)-1
    for left < right && right-left > 1 {
        mid := left + (right-left)/2
        if target <= nums[left] {
            return left
        } else if target > nums[left] && target < nums[mid] {
            right = mid-1
        } else if nums[mid] == target {
            return mid
        } else if target > nums[mid] && target < nums[right]  {
            left = mid+1
        } else if target > nums[right] {
            return right+1
        }
    }

    return left+1
}

func GF_searchInsert(nums []int, target int) int {
    n := len(nums)
    left, right := 0, n - 1
    ans := n
    for left <= right {
        mid := (right - left) >> 1 + left
        if target <= nums[mid] {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}

