package leetcode

import "golang-acm/util"

func levelOrder(root *util.MultiTreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    res := [][]int{}
    count := 1
    nodeArr := []*util.MultiTreeNode{root}
    for len(nodeArr) > 0 {
        nextCount := 0
        temp := []int{}
        for i := count; i > 0; i-- {
            node := nodeArr[0]
            temp = append(temp, node.Val)
            nodeArr = nodeArr[1:]
            nextCount += len(node.Children)
            nodeArr = append(nodeArr, node.Children...)
        }
        count = nextCount
        res = append(res, append([]int{}, temp...))
    }
    return res
}

func levelOrder2(root *util.MultiTreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    res := [][]int{}
    nodeArr := []*util.MultiTreeNode{root}
    for len(nodeArr) > 0 {
        temp := []int{}
        level := nodeArr
        nodeArr = nil
        for _, node := range level {
            temp = append(temp, node.Val)
            nodeArr = append(nodeArr, node.Children...)
        }
        res = append(res, append([]int{}, temp...))
    }
    return res
}

