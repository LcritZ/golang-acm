package leetcode

import "golang-acm/util"

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }

    dp := make([]int, len(nums))
    dp[0] = nums[0]
    dp[1] = util.Max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        dp[i] = util.Max((nums[i]+dp[i-2]), dp[i-1])
    }
    return dp[len(nums)-1]
}

func rob2(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }

    x := nums[0]
    y := util.Max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        x, y = y, util.Max((nums[i]+x), y)
    }
    return y
}
