package leetcode

import "fmt"

/**
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

 */
func BinarySearch(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    left, right := 0, len(nums)-1
    for left <= right {
        mid := left+(right-left)/2
        fmt.Println(mid, nums[mid], target)

        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid-1
        } else if nums[mid] < target {
            left = mid+1
        }
    }

    return -1
}

func BinarySearch2(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    left, right := 0, len(nums)
    //决定left right用什么区间，for循环条件就要对应起来，左右闭合: <=     左闭右开:<
    for left < right {
        mid := left+(right-left)/2
        fmt.Println(mid, nums[mid], target)
        if nums[mid] == target {
            //不同操作区别在这里体现
            right = mid
        } else if nums[mid] > target {
            right = mid
        } else if nums[mid] < target {
            left = mid+1
        }
    }

    //判断left未出界
    if left < len(nums) && target == nums[left] {
        return left
    } else {
        return -1
    }
}
