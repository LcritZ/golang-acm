package leetcode

import (
	"golang-acm/util"
)

func getPath(root, target *util.TreeNode) (path []*util.TreeNode) {
	node := root
	for node != target {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return
}

func lowestCommonAncestor(root, p, q *util.TreeNode) (ancestor *util.TreeNode) {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ancestor = pathP[i]
	}
	return
}

func lowestCommonAncestor2(root, p, q *util.TreeNode) *util.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}



//func lowestCommonAncestor2(root, p, q *util.TreeNode) *util.TreeNode {
//	ancestor := root
//	for {
//		if p.Val < ancestor.Val && q.Val < ancestor.Val {
//			ancestor = ancestor.Left
//		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
//			ancestor = ancestor.Right
//		} else {
//			return ancestor
//		}
//	}
//}
