package leetcode

import "golang-acm/util"

func IsSymmetric(root *util.TreeNode) bool {
	if root == nil {
		return true
	}
	return isNodeSymmetric(root, root)
}

func isNodeSymmetric(left, right *util.TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil {
		if left.Val == right.Val {
			return isNodeSymmetric(left.Left, right.Right) && isNodeSymmetric(left.Right, right.Left)
		} else {
			return false
		}
	} else {
		return false
	}

}
