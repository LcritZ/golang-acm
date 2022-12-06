package leetcode

import (
    "fmt"
    "golang-acm/util"
)

func ConstructFromPrePost(preorder []int, postorder []int) *util.TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    fmt.Println(preorder)
    fmt.Println(postorder)
    return constructTree(preorder, postorder, 0,0,len(preorder))
}

func constructTree(preorder []int, postorder []int, i1, i2, length int) *util.TreeNode {
    if length == 0 {
        return nil
    }

    root := &util.TreeNode{
        Val: preorder[i1],
    }
    if length == 1 {
        return root
    }

    leftLen, rightLen := 0, 0
    for i := i2; i < length+i2-1; i++ {
        if postorder[i] == preorder[i1+1]  {
            leftLen = i-i2+1
            rightLen = length-leftLen-1
            break
        }
    }

    fmt.Println(i1, i2, leftLen,rightLen)
    root.Left = constructTree(preorder, postorder, i1+1, i2, leftLen)
    root.Right = constructTree(preorder, postorder, i1+leftLen+1, i2+leftLen, rightLen)

    return root
}

