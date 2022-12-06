package basic

import (
    "fmt"
    "golang-acm/util"
)

func PrintTree(root *util.TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }

    count := 1
    nextCount := 0
    temp := []int{}
    nodeArr := []*util.TreeNode{root}
    for len(nodeArr) > 0 {
        node := nodeArr[0]
        temp = append(temp, node.Val)
        count--
        nodeArr = nodeArr[1:]
        if node.Left != nil {
            nodeArr = append(nodeArr, node.Left)
            nextCount++
        }
        if node.Right != nil {
            nodeArr = append(nodeArr, node.Right)
            nextCount++
        }

        if count == 0 {
            curr := make([]int, len(temp))
            copy(curr, temp)
            res = append(res, curr)
            temp = []int{}
            count = nextCount
            nextCount = 0
        }
    }

    return res
}

func PrintTree2(root *util.TreeNode) {
    if root == nil {
        return
    }

    nodeArr := []*util.TreeNode{root}
    first := true
    flag := false
    for len(nodeArr) > 0 && (flag || first) {
        length := len(nodeArr)
        i := 0
        s := ""
        flag = false
        for ; (flag || first) && i < length; i++ {
            node := nodeArr[0]
            nodeArr = nodeArr[1:]
            if node != nil {
                s += fmt.Sprintf("%d, ", node.Val)
            } else {
                s += "#, "
            }
            if node != nil && node.Left != nil {
                nodeArr = append(nodeArr, node.Left)
                flag = true
            } else {
                nodeArr = append(nodeArr, nil)
            }
            if node != nil && node.Right != nil {
                nodeArr = append(nodeArr, node.Right)
                flag = true
            } else {
                nodeArr = append(nodeArr, nil)
            }
        }
        fmt.Println(s)
        first = flag
    }
}
