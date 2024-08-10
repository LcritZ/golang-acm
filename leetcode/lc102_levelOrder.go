package leetcode

import "golang-acm/util"

func LevelOrderBinaryTree(root *util.TreeNode) [][]int {
	var res = [][]int{}
	var curr = []*util.TreeNode{}
	var next = []*util.TreeNode{}

	if root != nil {
		curr = append(curr, root)
	}
	for len(curr) > 0 {
		temp := make([]int, len(curr))
		for i := 0; i < len(curr); i++ {
			temp[i] = curr[i].Val
			if curr[i].Left != nil {
				next = append(next, curr[i].Left)
			}
			if curr[i].Right != nil {
				next = append(next, curr[i].Right)
			}
		}
		res = append(res, temp)
		curr = make([]*util.TreeNode, len(next))
		copy(curr, next)
		next = []*util.TreeNode{}
	}

	return res
}
