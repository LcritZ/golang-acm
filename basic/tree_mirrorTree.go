package basic

import "golang-acm/util"

func Mirror( pRoot *util.TreeNode) *util.TreeNode {
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

