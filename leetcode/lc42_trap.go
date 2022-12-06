package leetcode

import "golang-acm/util"

func GF_trap(height []int) (ans int) {
    n := len(height)
    if n == 0 {
        return
    }

    leftMax := make([]int, n)
    leftMax[0] = height[0]
    for i := 1; i < n; i++ {
        leftMax[i] = util.Max(leftMax[i-1], height[i])
    }

    rightMax := make([]int, n)
    rightMax[n-1] = height[n-1]
    for i := n - 2; i >= 0; i-- {
        rightMax[i] = util.Max(rightMax[i+1], height[i])
    }

    for i, h := range height {
        ans += util.Min(leftMax[i], rightMax[i]) - h
    }
    return
}

func trap2(height []int) (ans int) {
    if len(height) == 0 {
        return
    }

    left, right := 0, len(height)-1
    leftMax, rightMax := 0, 0
    for left < right {
        leftMax = util.Max(leftMax, height[left])
        rightMax = util.Max(rightMax, height[right])
        if height[left] <= height[right] {
            ans += leftMax-height[left]
            left++
        } else {
            ans += rightMax-height[right]
            right--
        }
    }

    return
}
