package acm

import "golang-acm/common"

/**
完全二叉树 所有节点数
var hight

4  3满的  2*n-1

n层
n-1层  M个节点
2*(n-1) -1 + 2*(M-1) +  M的左或空

K个节点 P个
K-P

node.Left


       1
    3     4
  2   3  6

 */


func getNodeNum(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	arr := []*common.TreeNode{root}

	for len(arr) != 0 {
		node := arr[0]
		ans++
		if node.Left != nil {
			arr = append(arr, node.Left)
		}
		if node.Right != nil {
			arr = append(arr, node.Right)
		}
		if len(arr) > 1 {
			arr = arr[1:]
		} else {
			break
		}
	}
	return ans
}