package leetcode

import "golang-acm/util"

/**
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数

思路: 中序遍历的思想，每遍历一个节点就+1
注意全局变量的设置
 */

var count = 0
var kth = -1

func kthSmallest(root *util.TreeNode, k int) int {
    count = 0
    kth = -1
    inorderTree(root, k)
    return kth
}

func inorderTree(root *util.TreeNode, k int) (bool) {
    if root == nil {
        return false
    }
    left := inorderTree(root.Left, k)
    if left {
        return true
    }
    count++
    if count == k {
        kth = root.Val
        return true
    }
    right := inorderTree(root.Right, k)
    if right {
        return true
    }

    return false
}
