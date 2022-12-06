package leetcode

import "golang-acm/util"

/**
递归的处理左右子树

 */
func constructMaximumBinaryTree(nums []int) *util.TreeNode {

    if len(nums) == 0 {
        return nil
    }

    max, index := nums[0], 0

    for i := 1; i < len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
            index = i
        }
    }

    root := &util.TreeNode{
        Val: max,
    }

    root.Left = constructMaximumBinaryTree(nums[0:index])
    if index == len(nums)-1 {
        root.Right = nil
    } else {
        root.Right = constructMaximumBinaryTree(nums[index+1:])
    }

    return root
}
