package basic

import (
	"fmt"
	"golang-acm/util"
)

/**
重建二叉树

 */
func ReConstructBinaryTree( pre []int ,  vin []int ) *util.TreeNode {
	// write code here
	if len(pre) == 0 || len(vin) == 0 || (len(pre) != len(vin)){
		var root util.TreeNode
		return &root
	}
	if len(pre) == 1 {
		var root = util.TreeNode{pre[0], nil,nil}
		return &root
	}
	rootVal := pre[0]
	temp := 0
	for i := 0; i< len(vin); i++ {
		if vin[i] == rootVal {
			temp = i
			break
		}
	}
	var root = util.TreeNode{rootVal, nil, nil}
	root.Left = ReConstructBinaryTree(pre[1:temp+1], vin[:temp])
	root.Right = ReConstructBinaryTree(pre[temp+1:], vin[temp+1:])
	return &root
}

func buildTree(pre []int, vin []int) *util.TreeNode {
	if len(pre) == 0 || len(vin) == 0 || (len(pre) != len(vin)){
		var root util.TreeNode
		return &root
	}
	if len(pre) == 1 {
		var root = util.TreeNode{pre[0], nil,nil}
		return &root
	}
	rootVal := pre[0]
	temp := 0
	for i := 0; i< len(vin); i++ {
		if vin[i] == rootVal {
			temp = i
			break
		}
	}
	var root = util.TreeNode{rootVal, nil, nil}
	root.Left = buildTree(pre[1:temp+1], vin[:temp])
	root.Right = buildTree(pre[temp+1:], vin[temp+1:])
	return &root
}

//通过代码
func BuildTree2(preorder []int, inorder []int) *util.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &util.TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	fmt.Println(i+1, len(inorder[:i])+1, "--")
	//root.Left = BuildTree2(preorder[1:len(inorder[:i])+1], inorder[:i])
	//root.Right = BuildTree2(preorder[len(inorder[:i])+1:], inorder[i+1:])
	root.Left = BuildTree2(preorder[1:i+1], inorder[:i])
	root.Right = BuildTree2(preorder[i+1:], inorder[i+1:])
	return root
}
