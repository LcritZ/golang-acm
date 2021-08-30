package acm

import (
	"fmt"
	"golang-acm/common"
)

var nodeArr []int

func DFS(pRoot *common.TreeNode)  {
	if pRoot == nil {
		return
	}
	if pRoot.Left != nil {
		DFS(pRoot.Left)
	}
	nodeArr = append(nodeArr, pRoot.Val)
	fmt.Println(nodeArr)
	if pRoot.Right != nil {
		DFS(pRoot.Right)
	}
}
