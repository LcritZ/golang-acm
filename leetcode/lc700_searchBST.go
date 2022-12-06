package leetcode

import (
    "golang-acm/util"
)

func searchBST(root *util.TreeNode, val int) *util.TreeNode {

    if root == nil {
        return nil
    }

    if root.Val == val {
        return root
    } else if root.Val > val {
        return searchBST(root.Left, val)
    } else {
        return searchBST(root.Right, val)
    }
}

func searchBST2(root *util.TreeNode, val int) *util.TreeNode {

    if root == nil {
        return nil
    }

    node := root
    for node != nil {
        if node.Val == val {
            return node
        } else if node.Val > val {
            node = node.Left
        } else {
            node = node.Right
        }
    }

    return node
}
