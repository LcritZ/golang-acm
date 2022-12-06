package leetcode

import "golang-acm/util"

func MaxRotateFunction(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }

    maxRes := 0
    for i := 0; i < len(nums); i++ {
        maxRes += i*nums[i]
    }

    curr := maxRes
    for j := 1; j < len(nums); j++ {
        curr = curr+sum - len(nums)*nums[len(nums)-j]
        maxRes = util.Max(maxRes, curr)
    }

    return maxRes
}

