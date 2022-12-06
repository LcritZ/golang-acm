package leetcode

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

思路：回溯的思想，取出来加进去，再回退

注意回溯的数据变更，不能影响后面的使用，比如left+1，下面应该使用之前的left, 只有传进去的参数+1

 */

func GenerateParenthesis(n int) []string {

    if n == 0 {
        return []string{}
    }

    res := []string{}

    var backtrack func(curr string, left, right int)
    backtrack = func(curr string, left, right int) {
        if len(curr) == 2*n {
            temp := curr
            res = append(res, temp)
        }

        if left < n {
            curr += "("
            backtrack(curr, left+1, right)
            curr = curr[0:len(curr)-1]
        }
        if right < left {
            curr += ")"
            backtrack(curr, left, right+1)
            curr = curr[0:len(curr)-1]
        }
    }

    backtrack("", 0, 0)

    return res
}
