package leetcode

import (
    "fmt"
    "strconv"
)

/**
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
n = 1 1
n = 2 2
n = 3 5
n = 4 14

选一个数作为根节点
左子树可能的情况 * 右子树的情况
累加起来

 */

var countMap = make(map[string]int) 

func NumTrees(n int) int {
    if n == 1 {
        return 1
    }
    
    return countTrees(1, n)
}

func countTrees(start, end int) int {
    if start >= end {
        return 1
    }
    res := 0
    left, right := 1, 1
    for i := start; i <= end; i++ {
        leftStr := fmt.Sprintf("%s-%s", strconv.Itoa(start), strconv.Itoa(i-1))
        if countMap[leftStr] > 0 {
            left = countMap[leftStr]
        } else {
            left = countTrees(start, i-1)
            countMap[leftStr] = left
        }
        rightStr := fmt.Sprintf("%s-%s", strconv.Itoa(i+1), strconv.Itoa(end))
        if countMap[rightStr] > 0 {
            right = countMap[rightStr]
        } else {
            right = countTrees(i+1, end)
            countMap[rightStr] = right
        }
        res += left*right
    }
    
    return res
}

func GF_numTrees(n int) int {
    C := 1
    for i := 0; i < n; i++ {
        C = C * 2 * (2 * i + 1) / (i + 2);
    }
    return C
}
