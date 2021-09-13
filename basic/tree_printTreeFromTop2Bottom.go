package basic

import "golang-acm/util"

func PrintFromTopToBottom( root *util.TreeNode) []int {
	var queue []*util.TreeNode
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
