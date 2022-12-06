package leetcode

import "fmt"

/***
给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。
输入：nums = [1,1,1], k = 2
输出：2

思路：通过前缀和，在顺序计算前缀和的时候，就判断在前面有多少个sum[i]-k的数存在，有多少结果就加多少

注意map需要保存, 注意获取map时需要判断是否存在
 */
func SubarraySum(nums []int, k int) int {
    res, sum := 0, 0
    numMap := make(map[int]int)
    numMap[0] = 1
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if _, ok := numMap[sum-k]; ok {
            res += numMap[sum-k]
        }
        numMap[sum]++
    }
    return res
}

func SubarraySum2(nums []int, k int) int {
    if len(nums) == 0 {
        return 0
    }

    sumMap := make(map[int]int)
    sum := 0
    res := 0
    //sumMap[0] = 1
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        sumMap[sum]++
        fmt.Println(sumMap)
        res += sumMap[sum-k]
    }

    return res
}

/**
map[0:1 1:1]
map[0:1 1:1 2:1]
map[0:1 1:1 2:1 3:1]
 */