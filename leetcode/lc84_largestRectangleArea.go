package leetcode

import "golang-acm/util"

/**
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

 */
func largestRectangleArea(heights []int) int {
    n := len(heights)
    //left, right保存一个高度能到达的左右下标
    left, right := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        right[i] = n
    }
    stack := []int{}
    for i := 0; i < n; i++ {
        for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
            right[stack[len(stack)-1]] = i
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            left[i] = -1
        } else {
            left[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }
    ans := 0
    for i := 0; i < n; i++ {
        ans = util.Max(ans, (right[i] - left[i] - 1) * heights[i])
    }
    return ans
}

