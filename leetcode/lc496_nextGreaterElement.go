package leetcode

import "fmt"

/**
1 暴力遍历

2 栈存储从后到前的数据

思路：map存储位置，

或者用栈的思想

 */
func NextGreaterElement(nums1, nums2 []int) []int {
    mp := map[int]int{}
    stack := []int{}
    for i := len(nums2) - 1; i >= 0; i-- {
        num := nums2[i]
        for len(stack) > 0 && num >= stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        }
        fmt.Println(stack)
        if len(stack) > 0 {
            mp[num] = stack[len(stack)-1]
        } else {
            mp[num] = -1
        }
        stack = append(stack, num)
    }
    res := make([]int, len(nums1))
    for i, num := range nums1 {
        res[i] = mp[num]
    }
    return res
}

