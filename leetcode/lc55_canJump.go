package leetcode

import "golang-acm/util"

/**
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

思路：就是找到0的位置，然后判断前一个位置是不是大于1，连续0的个数，


 */
func canJump(nums []int) bool {

    if len(nums) == 0 || nums[0] == 0 {
        return false
    }

    max := 0
    i := 0
    for ; i < len(nums); i++ {
        if nums[i] == 0 {
            max--
            if max == 0 {
                return false
            }
        } else {
            max = util.Max(max-1, nums[i])
        }
    }

    if i == len(nums)-1 {
        return true
    } else {
        return false
    }

}

//更新能走到最远的位置
func canJump2(nums []int) bool {

    next := 0
    for i := 0; i < len(nums); i++ {
        if i > next {
            return false
        }
        next = util.Max(next, i+nums[i])
    }

    return true
}