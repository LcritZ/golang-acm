package leetcode

import (
    "fmt"
)

/**
注意題目是否需要保持順序
 */
func MoveZeroes(nums []int)  {
    if len(nums) == 0 {
        return
    }

    slow, fast := 0, 0
    for fast < len(nums) {
        if nums[fast] != 0 {
            nums[slow], nums[fast] = nums[fast], nums[slow]
            slow++
        }
        fast++
    }
    fmt.Println(nums)
    return
}




func MoveZeroes1(nums []int)  {
    if len(nums) == 0 {
        return
    }

    left, right := 0, len(nums)-1
    for left < right {
        if nums[right] == 0 {
            right--
            continue
        }
        if nums[left] == 0 {
            temp := nums[right]
            nums[left] = temp
            nums[right] = 0
            right--
        }
        left++
    }

    fmt.Println(nums)

    return
}

