package basic

import "golang-acm/util"

func HasSubtree( pRoot1 *util.TreeNode,  pRoot2 *util.TreeNode) bool {
	// write code here
	if pRoot1 == nil || pRoot2 == nil{
		return false
	}
	var result bool = false
	if pRoot1.Val == pRoot2.Val {
		 result = JudgeNode(pRoot1, pRoot2)
	}
	if !result {
		result = HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
	}
	return result
}

func JudgeNode(root1 *util.TreeNode, root2 *util.TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return root1.Val == root2.Val && JudgeNode(root1.Left, root2.Left) && JudgeNode(root1.Right, root2.Right)
}