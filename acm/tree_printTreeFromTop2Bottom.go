package acm

import "golang-acm/common"

func PrintFromTopToBottom( root *common.TreeNode) []int {
	var queue []*common.TreeNode
	var res []int
	if root == nil {
		return res
	} else {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		res = append(res, queue[0].Val)
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}

	return res
}
