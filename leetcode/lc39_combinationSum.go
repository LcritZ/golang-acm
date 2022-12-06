package leetcode

var res39 [][]int

/**
回溯的重点
1 边界，遍历完成index==len就要退出
2 长度符合就填res
3 回溯需要判断是否跳过当前节点
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。


 */
func CombinationSum(candidates []int, target int) [][]int {

    combineHelper(candidates, []int{}, target, 0)
    return res39
}

func combineHelper(candidates []int, curr []int, target int, index int) {
    if index == len(candidates) {
        return
    }
    if target == 0 {
        temp := make([]int, len(curr))
        copy(temp, curr)
        res39 = append(res39, temp)
        return
    }



    combineHelper(candidates, curr, target, index+1)

    if target-candidates[index] >= 0 {
        curr = append(curr, candidates[index])
        combineHelper(candidates, curr, target-candidates[index], index)
        curr = curr[0:len(curr)-1]
    }

    return
}

func CombinationSum2(candidates []int, target int) [][]int {
    res := [][]int{}

    if len(candidates) == 0  {
        return res
    }

    curr := []int{}
    var backtrack func(index, target int)
    backtrack = func(index, target int) {
        if index == len(candidates) {
            return
        }
        if target == 0 {
            res = append(res, append([]int(nil), curr...))
            return
        }

        backtrack(index+1, target)

        if target-candidates[index] >= 0 {
            curr = append(curr, candidates[index])
            backtrack(index, target-candidates[index])
            curr = curr[:len(curr)-1]
        }
    }

    backtrack(0, target)
    return res
}

func GF_combinationSum(candidates []int, target int) (ans [][]int) {
    comb := []int{}
    var dfs func(target, idx int)
    dfs = func(target, idx int) {
        if idx == len(candidates) {
            return
        }
        if target == 0 {
            ans = append(ans, append([]int(nil), comb...))
            return
        }
        // 直接跳过
        dfs(target, idx+1)
        // 选择当前数
        if target-candidates[idx] >= 0 {
            comb = append(comb, candidates[idx])
            dfs(target-candidates[idx], idx)
            comb = comb[:len(comb)-1]
        }
    }
    dfs(target, 0)
    return
}
