package leetcode

import (
	"golang-acm/util"
)


func maxDepth(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth+1
	} else {
		return rightDepth+1
	}
}
