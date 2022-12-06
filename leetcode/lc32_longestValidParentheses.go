package leetcode

import "golang-acm/util"

/**
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

*/
func LongestValidParentheses(s string) int {
    if len(s) == 0 {
        return 0
    }
    max := 0
    left, right := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if right == left {
            max = util.Max(max, 2*right)
        } else if right > left {
            left, right = 0, 0
        }
    }

    left, right = 0, 0
    for i := len(s)-1; i >=0; i-- {
        if s[i] == ')' {
            right++
        } else {
            left++
        }

        if right == left {
            max = util.Max(max, 2*right)
        } else if left > right {
            left, right = 0, 0
        }
    }

    return max
}

