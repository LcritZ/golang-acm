package leetcode

import "golang-acm/util"

func buildTree106(inorder []int, postorder []int) *util.TreeNode {

    if len(inorder) == 0 {
        return nil
    }

    root := &util.TreeNode{
        Val: postorder[len(postorder)-1],
    }

    index := 0
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == root.Val {
            index = i
            break
        }
    }

    if len(postorder) > 1 {
        root.Left = buildTree106(inorder[0:index], postorder[0:index])
        root.Right = buildTree106(inorder[index+1:], postorder[index:len(postorder)-1])
    }
    return root

}
