package leetcode

import "golang-acm/util"

/*
插入节点，保持二插搜索
最好的一种情况就是插入的点放在叶子节点

 */
func insertIntoBST(root *util.TreeNode, val int) *util.TreeNode {
    if root == nil {
        return &util.TreeNode{
            Val: val,
        }
    }

    if val > root.Val {
        root.Right = insertIntoBST(root.Right, val)
    } else {
        root.Left = insertIntoBST(root.Left, val)
    }
    return root
}

func GF_insertIntoBST(root *util.TreeNode, val int) *util.TreeNode {
    if root == nil {
        return &util.TreeNode{Val: val}
    }
    p := root
    for p != nil {
        if val < p.Val {
            if p.Left == nil {
                p.Left = &util.TreeNode{Val: val}
                break
            }
            p = p.Left
        } else {
            if p.Right == nil {
                p.Right = &util.TreeNode{Val: val}
                break
            }
            p = p.Right
        }
    }
    return root
}
