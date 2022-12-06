package leetcode

import "golang-acm/util"

var preorderArr []int

func preorderTraversal(root *util.TreeNode) []int {
    if root == nil {
        return []int{}
    }

    preorderArr = []int{}
    preorderTree(root)
    return preorderArr
}

func preorderTree(root *util.TreeNode) {
    if root == nil {
        return
    }
    preorderArr = append(preorderArr, root.Val)
    preorderTree(root.Left)
    preorderTree(root.Right)
}

//迭代的方式，把left全部加进去之后，再加right
func GF_preorderTraversal(root *util.TreeNode) (vals []int) {
    stack := []*util.TreeNode{}
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
            vals = append(vals, node.Val)
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1].Right
        stack = stack[:len(stack)-1]
    }
    return
}
