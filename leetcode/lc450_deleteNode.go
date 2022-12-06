package leetcode

import (
    "golang-acm/util"
)

/**
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。

        5
    3       10
 2    4   6    11
            8
           7  9
        6
    3       10
 2    4   8   11
         7 9


        5
    3       10
 2    4   8    11
         7  9
        6

        6
    3       10
 2    4   8   11
         7 9

思路：删除该节点，需要用该节点的中序遍历下一个节点代替，也就是右子树的最左边节点

代替完之后，需要将最左节点的右节点补充上去



 */

func getSuccessor450(node *util.TreeNode) *util.TreeNode {
    succ := node.Right
    for succ.Left != nil {
        succ = succ.Left
    }
    return succ
}

func predecessor(node *util.TreeNode) *util.TreeNode {
    pred := node.Left
    for pred.Right != nil {
        pred = pred.Right
    }
    return pred
}

func deleteNode(root *util.TreeNode, key int) *util.TreeNode {
    if root == nil {
        return nil
    }
    if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else {
        if root.Right == nil && root.Left == nil {
            root = nil
            return root
        } else if root.Right != nil {
            root.Val = getSuccessor450(root).Val
            root.Right = deleteNode(root.Right, root.Val)
        } else {
            root.Val = predecessor(root).Val
            root.Left = deleteNode(root.Left, root.Val)
        }
    }
    return root
}
