package leetcode

import "golang-acm/util"

func RightSideView(root *util.TreeNode) []int {
    res := []int{}

    if root == nil {
        return res
    }

    nodeArr := []*util.TreeNode{root}

    for len(nodeArr) > 0 {
        count := len(nodeArr)
        for i := 0; i < count; i++ {
            node := nodeArr[0]
            nodeArr = nodeArr[1:]
            if i == count-1 {
                res = append(res, node.Val)
            }

            if node.Left != nil {
                nodeArr = append(nodeArr, node.Left)
            }
            if node.Right != nil {
                nodeArr = append(nodeArr, node.Right)
            }
        }
    }

    return res
}

func RightSideViewDfs(root *util.TreeNode) []int {

    res := []int{}

    var dfs func(root *util.TreeNode, level int)
    dfs = func(root *util.TreeNode, level int) {
        if root == nil {
            return
        }
        if level > len(res) {
            res = append(res, root.Val)
        }
        dfs(root.Right,level+1)
        dfs(root.Left, level+1)
    }
    dfs(root, 1)
    return res
}
