package leetcode

import "golang-acm/util"

/**
思路：中序遍历，然后将数据填入栈，出栈的时候统计大于等于自己的数据

 */

var nodeStack = []*util.TreeNode{}

func convertBST(root *util.TreeNode) *util.TreeNode {
    if root == nil {
        return nil
    }
    inorderTree538(root)
    length := len(nodeStack)
    count := 0
    for i := length-1; i >= 0; i-- {
        count += nodeStack[i].Val
        nodeStack[i].Val = count
    }
    return root
}

func inorderTree538(root *util.TreeNode) {
    if root == nil {
        return
    }

    inorderTree538(root.Left)
    nodeStack = append(nodeStack, root)
    inorderTree538(root.Right)
}

//遍历过程中就加上对应的值
func GF_convertBST(root *util.TreeNode) *util.TreeNode {
    sum := 0
    var dfs func(*util.TreeNode)
    dfs = func(node *util.TreeNode) {
        if node != nil {
            dfs(node.Right)
            sum += node.Val
            node.Val = sum
            dfs(node.Left)
        }
    }
    dfs(root)
    return root
}

func getSuccessor(node *util.TreeNode) *util.TreeNode {
    succ := node.Right
    for succ.Left != nil && succ.Left != node {
        succ = succ.Left
    }
    return succ
}

func GF_convertBST2(root *util.TreeNode) *util.TreeNode {
    sum := 0
    node := root
    for node != nil {
        if node.Right == nil {
            sum += node.Val
            node.Val = sum
            node = node.Left
        } else {
            succ := getSuccessor(node)
            if succ.Left == nil {
                succ.Left = node
                node = node.Right
            } else {
                succ.Left = nil
                sum += node.Val
                node.Val = sum
                node = node.Left
            }
        }
    }
    return root
}
