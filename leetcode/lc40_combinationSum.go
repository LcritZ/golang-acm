package leetcode

import (
    "fmt"
    "sort"
)

/**
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。


细节在于遍历下一个需要跳过数值一样的节点，不然出现重复的
然后取当前节点，这不能跳过，需要使用数值一样的的

 */
func CombinationSum40(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    fmt.Println(candidates)
    res := [][]int{}
    if len(candidates) == 0 {
        return res
    }

    curr := []int{}
    var backtrack func(target, index int)
    backtrack = func(target, index int) {
        if target == 0 {
            temp := make([]int, len(curr))
            copy(temp, curr)
            res = append(res, temp)
            return
        }
        if index == len(candidates) {
            return
        }

        next := index+1
        for next < len(candidates) && candidates[next] == candidates[next-1]  {
            next++
        }
        backtrack(target, next)

        if target-candidates[index] >= 0 {
            curr = append(curr, candidates[index])
            backtrack(target-candidates[index], index+1)
            curr = curr[:len(curr)-1]
        }
    }

    backtrack(target, 0)

    return res
}
