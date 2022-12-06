package leetcode

import (
    "golang-acm/util"
    "strconv"
)

/**
给定一棵二叉树 root，返回所有重复的子树。

对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

如果两棵树具有相同的结构和相同的结点值，则它们是重复的

思路：


 */


var DuplicateRes = []*util.TreeNode{}
var NodeMap = make(map[string]int)

func findDuplicateSubtrees(root *util.TreeNode) []*util.TreeNode {
    //注意需要更新全局变量
    DuplicateRes = []*util.TreeNode{}
    NodeMap = make(map[string]int)
    serializedTree(root)
    return DuplicateRes
}

func serializedTree(root *util.TreeNode) string {
    if root == nil {
        return "#"
    }

    left := serializedTree(root.Left)
    right := serializedTree(root.Right)

    curr := strconv.Itoa(root.Val)+ "-"+ left + "-" + right
    NodeMap[curr]++
    //等于2就加，大于2已经加过了，不用重复加
    if NodeMap[curr] == 2 {
        DuplicateRes = append(DuplicateRes, root)
    }

    return curr
}
