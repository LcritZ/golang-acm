package basic

import "golang-acm/util"

func inorderTraversal( root *util.TreeNode ) []int {
    // write code here
    res := []int{}
    if root == nil {
        return res
    }

    nodeArr := []*util.TreeNode{}
    nodeArr = append(nodeArr, root)

    nodeMap := make(map[int]bool)

    for len(nodeArr) > 0 {
        node := nodeArr[len(nodeArr)-1]
        if node.Left != nil && !nodeMap[node.Left.Val] {
            nodeArr = append(nodeArr, node.Left)
            continue
        }
        res = append(res, node.Val)
        nodeArr = nodeArr[:len(nodeArr)-1]
        nodeMap[node.Val] = true
        if node.Right != nil {
            nodeArr = append(nodeArr, node.Right)
        }
    }

    return res
}


func inorderTraversal2( root *util.TreeNode ) []int {
    // write code here
    res := []int{}
    if root == nil {
        return res
    }

    nodeArr := []*util.TreeNode{}
    for root != nil || len(nodeArr) > 0 {
        for root != nil {
            nodeArr = append(nodeArr, root)
            root = root.Left
        }
        root = nodeArr[len(nodeArr)-1]
        res = append(res, root.Val)
        nodeArr = nodeArr[:len(nodeArr)-1]
        root = root.Right
    }

    return res
}

func HL_PrintTree(root *util.TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }

    nodeArr := []*util.TreeNode{root}
    flag := 1
    for len(nodeArr) > 0 {
        i := 0
        length := len(nodeArr)
        temp := make([]*util.TreeNode, length)
        copy(temp, nodeArr)
        for ; i < length; i++ {
            if flag%2 == 1 {
                res = append(res, temp[length-1-i].Val)
            } else {
                res = append(res, temp[i].Val)
            }
            node := nodeArr[0]
            nodeArr = nodeArr[1:]
            if node.Right != nil {
                nodeArr = append(nodeArr, node.Right)
            }
            if node.Left != nil {
                nodeArr = append(nodeArr, node.Left)
            }
        }
        flag++
    }

    return res
}
