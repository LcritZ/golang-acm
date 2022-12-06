package leetcode

import "golang-acm/util"

/**
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9

思路：递归赋值左右


 */

func InvertTree(root *util.TreeNode) *util.TreeNode {
    if root == nil {
        return nil
    }
    InvertTree(root.Left)
    InvertTree(root.Right)
    root.Left, root.Right = root.Right, root.Left
    return root
}
