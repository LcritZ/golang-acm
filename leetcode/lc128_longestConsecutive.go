package leetcode

import "golang-acm/util"

func LongestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    res := 0
    numMap := make(map[int]struct{})
    for i := 0; i < len(nums); i++ {
        numMap[nums[i]] = struct{}{}
    }

    for i := 0; i < len(nums); i++ {
        key := nums[i]
        if _, ok := numMap[key-1]; ok {
            continue
        }
        curr := 1
        _, ok := numMap[key+1]
        for ok {
             curr++
             key++
             _, ok = numMap[key+1]
        }
        res = util.Max(res, curr)
    }

    return res

}

