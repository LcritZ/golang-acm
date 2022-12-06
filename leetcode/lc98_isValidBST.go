package leetcode

import (
    "golang-acm/util"
)

func isValidBST(root *util.TreeNode) bool {
    if root == nil {
        return true
    }

    return isValidBSTHelper(root, nil, nil)

}

func isValidBSTHelper(root *util.TreeNode, min *util.TreeNode, max *util.TreeNode) bool {
    if root == nil {
        return true
    }
    if min != nil && root.Val <= min.Val {
        return false
    }
    if max != nil && root.Val >= max.Val {
        return false
    }

    return isValidBSTHelper(root.Left, min, root) && isValidBSTHelper(root.Right, root, max)
}
