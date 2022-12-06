package leetcode

import "golang-acm/util"

func jump(nums []int) int {
    length := len(nums)
    start, curr := 0, 0

    step := 0

    for i := 0; i < length-1; i++ {
        curr = util.Max(curr, i + nums[i])
        if i == start {
            start = curr
            step++
        }
    }

    return step
}
