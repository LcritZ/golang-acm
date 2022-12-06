package leetcode

import (
    "golang-acm/util"
)

var diameter int
func diameterOfBinaryTree(root *util.TreeNode) int {
    if root == nil {
        return 0
    }
    diameter = 0
    postorderTree(root)
    return diameter
}


func postorderTree(root *util.TreeNode) int {
    if root == nil {
        return 0
    }
    left := postorderTree(root.Left)
    right := postorderTree(root.Right)

    curr := util.Max(left, right)
    curr++
    diameter = util.Max(left+right+1, diameter)

    return curr
}
