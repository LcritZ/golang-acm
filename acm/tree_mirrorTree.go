package acm

import "golang-acm/common"

func Mirror( pRoot *common.TreeNode) *common.TreeNode {
	// write code here
	//var nodeArr []int
	if pRoot == nil {
		return nil
	}
	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}

