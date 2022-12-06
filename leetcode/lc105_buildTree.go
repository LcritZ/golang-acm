package leetcode

import "golang-acm/util"

/**
前序和中序遍历

 */
func buildTree105(preorder []int, inorder []int) *util.TreeNode {

    if len(preorder) != len(inorder) || len(preorder) == 0 {
        return nil
    }

    root := &util.TreeNode{Val: preorder[0]}
    index := 0
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            index = i
            break
        }
    }


    root.Left = buildTree105(preorder[1:1+index], inorder[0:index])
    root.Right = buildTree105(preorder[1+index:], inorder[index+1:])


    return root
}

func buildTree105_2(preorder []int, inorder []int) *util.TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &util.TreeNode{preorder[0], nil, nil}
    stack := []*util.TreeNode{}
    stack = append(stack, root)
    var inorderIndex int
    for i := 1; i < len(preorder); i++ {
        preorderVal := preorder[i]
        node := stack[len(stack)-1]
        if node.Val != inorder[inorderIndex] {
            node.Left = &util.TreeNode{preorderVal, nil, nil}
            stack = append(stack, node.Left)
        } else {
            for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
                node = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                inorderIndex++
            }
            node.Right = &util.TreeNode{preorderVal, nil, nil}
            stack = append(stack, node.Right)
        }
    }
    return root
}
