package leetcode

import (
    "fmt"
    "golang-acm/util"
)

/**
思路：按层遍历，然后处理是否是每一层的最后一个节点，
标准通过level每一层节点判断，到下一层，更新level总节点，更新计数器

 */
func Connect(root *util.SpecialTreeNode) *util.SpecialTreeNode {
    if root == nil {
        return nil
    }

    nodeQueue := []*util.SpecialTreeNode{root}
    count := 1
    level := 1
    for len(nodeQueue) > 0 {
        node := nodeQueue[0]
        nodeQueue = nodeQueue[1:]
        if count == level {
            count = 0
            level = 2*level
            node.Next = nil
        } else {
            if len(nodeQueue) > 0 {
                node.Next = nodeQueue[0]
            } else {
                node.Next = nil
            }
        }
        count++
        if node.Left != nil {
            nodeQueue = append(nodeQueue, node.Left, node.Right)
        }
        fmt.Println(len(nodeQueue))
    }
    return root
}

/**
按层处理，每一层把下一层的next指向确定，通过当前比下一层高，当前已经用next连起来了，下一层就可以关联起来，
 */
func GF_connect(root *util.SpecialTreeNode) *util.SpecialTreeNode {
    if root == nil {
        return root
    }

    // 每次循环从该层的最左侧节点开始
    for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
        // 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
        for node := leftmost; node != nil; node = node.Next {
            // 左节点指向右节点
            node.Left.Next = node.Right

            // 右节点指向下一个左节点
            if node.Next != nil {
                node.Right.Next = node.Next.Left
            }
        }
    }

    // 返回根节点
    return root
}

/**
递归写法

 */
func GF_connect2(root *util.SpecialTreeNode) *util.SpecialTreeNode {
    if root == nil {
        return nil
    }

    connectNode(root.Left, root.Right)
    return root

}

func connectNode(node1 *util.SpecialTreeNode, node2 *util.SpecialTreeNode) {
    if node1 == nil || node2 == nil {
        return
    }

    node1.Next = node2
    connectNode(node1.Left, node1.Right)
    connectNode(node2.Left, node2.Right)
    connectNode(node1.Right, node2.Left)
    return
}