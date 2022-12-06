package leetcode

import (
	"golang-acm/util"
)

func minDepth(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	depth :=  bfs111(root)
	return depth
}

func bfs111(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	nodeArr := []*util.TreeNode{root}
	depth := 1
	num := 1
	for len(nodeArr) != 0 {
		num = len(nodeArr)
		//通过遍历的长度约束这一层节点，每次更新长度，就更新了这一层
		for i := 0; i < num; i++ {
			node := nodeArr[0]
			nodeArr = nodeArr[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				nodeArr = append(nodeArr, node.Left)
			}
			if node.Right != nil {
				nodeArr = append(nodeArr, node.Right)
			}
		}
		depth++
	}

	return depth
}
