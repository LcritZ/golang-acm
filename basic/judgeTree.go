package basic

import "golang-acm/util"

/**
 3           root
 2	  node1        node2
 1 leftH  rightH

判断一棵树是否是平衡树

*/

func JudgeTree(root *util.TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftRes, leftH := JudgeTree(root.Left)
	rightRes, rightH := JudgeTree(root.Right)
	if leftRes && rightRes {
		if (leftH-rightH) == 1 || (leftH-rightH) == 0 {
			return true, leftH+1
		} else if leftH-rightH == -1 {
			return true, rightH+1
		} else {
			return false, 0
		}
	}
	return false, 0
}


