package basic

import (
	"fmt"
	"golang-acm/util"
)

var nodeArr []int

func DFS(pRoot *util.TreeNode)  {
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
